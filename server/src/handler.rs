use std::sync::Arc;
use std::sync::LazyLock;
use std::sync::RwLock;
use std::time::Duration;

use base64::engine::general_purpose;
use base64::Engine;
use log::info;
use log::warn;
use once_cell::sync::Lazy;
use regex::Regex;
use reqwest::header::HeaderMap;
use reqwest::header::HeaderValue;
use reqwest::Certificate;
use tokio::sync::RwLock as AsyncRwLock;
use tokio::time::interval;

use crate::process;
use crate::state::GameEvent;
use crate::state::GameState;
use crate::state::GAME_STATE_MACHINE;

pub struct LeagueProcessMonitor {
    pid: Arc<AsyncRwLock<Option<u32>>>, // 使用 RwLock 来保护可变状态
    interval_ms: u64,
}

impl LeagueProcessMonitor {
    pub async fn new(interval_ms: u64) -> Arc<Self> {
        let monitor = Arc::new(Self { pid: Arc::new(AsyncRwLock::new(None)), interval_ms });

        let monitor_clone = Arc::clone(&monitor);
        tokio::spawn(async move {
            monitor_clone.run().await;
        });

        monitor
    }

    pub async fn run(self: Arc<Self>) {
        let mut interval = interval(Duration::from_millis(self.interval_ms));

        loop {
            interval.tick().await;

            let pid_result = process::find_process_id_by_name(CLIENT_EXE);

            match pid_result {
                // 读取到游戏进程
                Ok(new_pid) => {
                    let mut pid_lock = self.pid.write().await;
                    match *pid_lock {
                        // 当PID从无到有时
                        None => {
                            *pid_lock = Some(new_pid);
                            info!("Process found with PID: {}", new_pid);
                            self.send_event(GameEvent::LaunchSuccess).await;
                            self.setup_client(new_pid).await;
                        }
                        // 当PID有且改变时
                        Some(current_pid) if current_pid != new_pid => {
                            *pid_lock = Some(new_pid);
                            info!("Process PID changed from {} to {}", current_pid, new_pid);
                            self.send_event(GameEvent::Reconnect).await;
                            self.setup_client(new_pid).await;
                        }
                        // PID未变化，无需处理
                        _ => {}
                    }
                }

                // 未读取到游戏进程
                Err(_) => {
                    let pid_lock = self.pid.read().await;
                    if pid_lock.is_some() {
                        if !GAME_STATE_MACHINE.read().await.get_state().eq(&GameState::Disconnected) {
                            self.send_event(GameEvent::Disconnect).await;
                        }
                        warn!("Process not found. PID {} is no longer active.", pid_lock.unwrap());
                    } else {
                        // 当PID无且未获取到时
                        warn!("Process launch failed, unable to find PID.");
                        self.send_event(GameEvent::LaunchFailed).await;
                    }
                }
            }
        }
    }

    async fn setup_client(&self, pid: u32) {
        if !process::is_admin() {
            process::elevated();
        }
        let cmdline = process::get_process_cmdline(pid).unwrap();
        info!("Setting up client with command line: {}", cmdline);

        let (path, cmdline) = cmdline.split_once(' ').unwrap();
        let port = Regex::new(PORT_RE).unwrap().captures(cmdline).unwrap()[1].to_string();
        let token = Regex::new(TOKEN_RE).unwrap().captures(cmdline).unwrap()[1].to_string();
        let platform = Regex::new(PLATFORM_RE).unwrap().captures(cmdline).unwrap()[1].to_string();

        println!("URL: http://127.0.0.1:{}", port);
        println!("BASIC AUTH: riot / {}", token);

        let mut headers = HeaderMap::new();
        let token_encoded = general_purpose::STANDARD.encode(format!("riot:{}", token));

        let basic = HeaderValue::from_str(format!("Basic {}", token_encoded).as_str()).unwrap();
        headers.insert("Authorization", basic);

        let client = reqwest::ClientBuilder::new()
            .add_root_certificate(CERT.clone())
            .default_headers(headers)
            .timeout(Duration::from_millis(500))
            .build()
            .unwrap();

        let mut client_lock = GLOBAL_CLIENT.write().unwrap();
        *client_lock = Some(client); // 设置全局客户端
    }

    async fn send_event(&self, event: GameEvent) {
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(event);
    }
}

// 全局存储 reqwest::Client
static GLOBAL_CLIENT: Lazy<Arc<RwLock<Option<reqwest::Client>>>> =
    Lazy::new(|| Arc::new(RwLock::new(None)));

const CLIENT_EXE: &str = "LeagueClientUx.exe";

const PORT_RE: &str = r"--app-port=(\d+)";
const TOKEN_RE: &str = r"--remoting-auth-token=(\w+)";
const PLATFORM_RE: &str = r"--rso_platform_id=(\w+)";

static CERT: LazyLock<Certificate> = LazyLock::new(|| {
    Certificate::from_pem(
        br#"-----BEGIN CERTIFICATE-----
MIIEIDCCAwgCCQDJC+QAdVx4UDANBgkqhkiG9w0BAQUFADCB0TELMAkGA1UEBhMC
VVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFTATBgNVBAcTDFNhbnRhIE1vbmljYTET
MBEGA1UEChMKUmlvdCBHYW1lczEdMBsGA1UECxMUTG9MIEdhbWUgRW5naW5lZXJp
bmcxMzAxBgNVBAMTKkxvTCBHYW1lIEVuZ2luZWVyaW5nIENlcnRpZmljYXRlIEF1
dGhvcml0eTEtMCsGCSqGSIb3DQEJARYeZ2FtZXRlY2hub2xvZ2llc0ByaW90Z2Ft
ZXMuY29tMB4XDTEzMTIwNDAwNDgzOVoXDTQzMTEyNzAwNDgzOVowgdExCzAJBgNV
BAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRUwEwYDVQQHEwxTYW50YSBNb25p
Y2ExEzARBgNVBAoTClJpb3QgR2FtZXMxHTAbBgNVBAsTFExvTCBHYW1lIEVuZ2lu
ZWVyaW5nMTMwMQYDVQQDEypMb0wgR2FtZSBFbmdpbmVlcmluZyBDZXJ0aWZpY2F0
ZSBBdXRob3JpdHkxLTArBgkqhkiG9w0BCQEWHmdhbWV0ZWNobm9sb2dpZXNAcmlv
dGdhbWVzLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKoJemF/
6PNG3GRJGbjzImTdOo1OJRDI7noRwJgDqkaJFkwv0X8aPUGbZSUzUO23cQcCgpYj
21ygzKu5dtCN2EcQVVpNtyPuM2V4eEGr1woodzALtufL3Nlyh6g5jKKuDIfeUBHv
JNyQf2h3Uha16lnrXmz9o9wsX/jf+jUAljBJqsMeACOpXfuZy+YKUCxSPOZaYTLC
y+0GQfiT431pJHBQlrXAUwzOmaJPQ7M6mLfsnpHibSkxUfMfHROaYCZ/sbWKl3lr
ZA9DbwaKKfS1Iw0ucAeDudyuqb4JntGU/W0aboKA0c3YB02mxAM4oDnqseuKV/CX
8SQAiaXnYotuNXMCAwEAATANBgkqhkiG9w0BAQUFAAOCAQEAf3KPmddqEqqC8iLs
lcd0euC4F5+USp9YsrZ3WuOzHqVxTtX3hR1scdlDXNvrsebQZUqwGdZGMS16ln3k
WObw7BbhU89tDNCN7Lt/IjT4MGRYRE+TmRc5EeIXxHkQ78bQqbmAI3GsW+7kJsoO
q3DdeE+M+BUJrhWorsAQCgUyZO166SAtKXKLIcxa+ddC49NvMQPJyzm3V+2b1roP
SvD2WV8gRYUnGmy/N0+u6ANq5EsbhZ548zZc+BI4upsWChTLyxt2RxR7+uGlS1+5
EcGfKZ+g024k/J32XP4hdho7WYAS2xMiV83CfLR/MNi8oSMaVQTdKD8cpgiWJk3L
XWehWA==
-----END CERTIFICATE-----"#,
    )
    .unwrap()
});
