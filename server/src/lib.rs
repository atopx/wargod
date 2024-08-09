use std::sync::Mutex;
use std::thread::sleep;
use std::time::Duration;

use tauri::Emitter;
use tauri::Wry;

pub mod error;
pub mod handler;
pub mod process;

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tauri::command]
fn greet(name: &str) -> String {
    match handler::LeagueClient::new() {
        Ok(client) => {
            format!("PATH: {}\nTOKEN: {}\nPORT: {}", client.path, client.token, client.port)
        }
        Err(e) => format!("get league client error: {}", e),
    }
}

struct GameState {
    status: Mutex<String>,
}

#[tauri::command]
fn mock_connected(state: tauri::State<GameState>, app_handle: tauri::AppHandle<Wry>) {
    // let status = state.status.lock().unwrap().clone();
    sleep(Duration::from_secs(1));
    app_handle.emit("game-status-change", "Connected").unwrap();
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    if !process::is_admin() {
        process::elevated()
    }

    tauri::Builder::default()
        .manage(GameState { status: Mutex::new("Connecting".to_string()) })
        .plugin(tauri_plugin_shell::init())
        // .setup(|app| {
        //     // 获取应用程序的窗口句柄
        //     let handle = app.handle().clone();
        //     // 在启动时模拟游戏启动状态
        //     tauri::async_runtime::spawn(async move {
        //         // 模拟3秒的加载时间
        //         sleep(Duration::from_secs(3)).await;
        //         // 发送连接中的状态
        //         // handle.emit("game-status-change", "Connecting").unwrap();
        //         // 再次延迟1秒后发送已连接状态
        //         // sleep(Duration::from_secs(1)).await;
        //         handle.emit("game-status-change", "Connected").unwrap();
        //     });
        //     Ok(())
        // })
        .invoke_handler(tauri::generate_handler![greet, mock_connected])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
