use std::sync::LazyLock;
use std::time::Duration;

use base64::engine::general_purpose;
use base64::Engine;
use regex::Regex;
use reqwest::header::HeaderMap;
use reqwest::header::HeaderValue;
use reqwest::Certificate;

use crate::error;
use crate::process;

#[derive(Default, Debug, Clone)]
pub struct LeagueClient {
    timeout: Option<u64>,
    pub path: String,
    pub token: String,
    pub port: String,
    pub platform: String,
    pub client: reqwest::Client,
}

impl LeagueClient {
    pub fn refresh(&mut self) -> Result<(), error::ProcessError> {
        match process::find_process_id_by_name(CLIENT_EXE) {
            Ok(pid) => match process::get_process_cmdline(pid) {
                Ok(cmdline) => {
                    self.reset(cmdline);
                    Ok(())
                }
                Err(e) => Err(e),
            },
            Err(e) => Err(e),
        }
    }

    pub fn new() -> Result<Self, error::ProcessError> {
        let mut client = Self::default();
        match client.refresh() {
            Ok(_) => Ok(client),
            Err(e) => Err(e),
        }
    }

    pub fn reset(&mut self, cmdline: String) {
        let (path, cmdline) = cmdline.split_once(' ').unwrap();
        self.path = path.to_string();
        self.port = Regex::new(PORT_RE).unwrap().captures(cmdline).unwrap()[1].to_string();
        self.token = Regex::new(TOKEN_RE).unwrap().captures(cmdline).unwrap()[1].to_string();
        self.platform = Regex::new(PLATFORM_RE).unwrap().captures(cmdline).unwrap()[1].to_string();
        let mut headers = HeaderMap::new();
        let token = general_purpose::STANDARD.encode(format!("riot:{}", self.token));
        let basic = HeaderValue::from_str(format!("Basic {}", token).as_str()).unwrap();
        headers.insert("Authorization", basic);
        self.client = reqwest::ClientBuilder::new()
            .add_root_certificate(CERT.clone())
            .default_headers(headers)
            .timeout(Duration::from_millis(self.timeout.unwrap_or(500)))
            .build()
            .unwrap();
    }

    /// 动态监测游戏进程是否存货或发生改变
    pub async fn keepalive() {
        // tokio oneshot + timer.tick
    }
}

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
