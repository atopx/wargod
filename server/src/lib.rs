pub mod error;
pub mod handler;
pub mod process;

// Learn more about Tauri commands at https://tauri.app/v1/guides/features/command
#[tauri::command]
fn greet(name: &str) -> String {
    let client = handler::LeagueClient::new().unwrap();
    format!("PATH: {}\nTOKEN: {}\nPORT: {}", client.path, client.token, client.port)
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    if !process::is_admin() {
        process::elevated()
    }

    tauri::Builder::default()
        .plugin(tauri_plugin_shell::init())
        .invoke_handler(tauri::generate_handler![greet])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
