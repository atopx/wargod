use std::sync::Mutex;
use std::thread::sleep;
use std::time::Duration;

use tauri::Emitter;
use tauri::Wry;

pub mod error;
pub mod handler;
pub mod process;
pub mod state;

struct AppState {
    status: Mutex<state::GameState>,
}

#[tauri::command]
fn mock_connected(state: tauri::State<AppState>, app_handle: tauri::AppHandle<Wry>) {
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
        .manage(AppState { status: Mutex::new(state::GameState::default()) })
        .plugin(tauri_plugin_shell::init())
        .setup(|app| {
            // 获取应用程序的窗口句柄
            // let handle = app.handle().clone();
            // 在启动时模拟游戏启动状态
            // tauri::async_runtime::spawn(async move {
            //     // 模拟3秒的加载时间
            //     sleep(Duration::from_secs(3));
            //     // 发送连接中的状态
            //     // sleep(Duration::from_secs(1)).await;
            //     handle.emit("game-status-change", "Connected").unwrap();
            // });
            Ok(())
        })
        .invoke_handler(tauri::generate_handler![mock_connected])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
