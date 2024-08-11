use std::sync::Arc;

use log::LevelFilter;
use tokio::task;
use tokio::time::Duration;
use wargod_lib::handler::LeagueProcessMonitor;

fn init_logging() { env_logger::Builder::from_default_env().filter_level(LevelFilter::Info).init(); }

#[tokio::main]
async fn main() {
    // 初始化日志系统
    init_logging();

    // 创建并启动 LeagueProcessMonitor
    let interval_ms = 1000; // 每秒检查一次
    let monitor = LeagueProcessMonitor::new(interval_ms).await;

    // 在后台运行监控器
    let monitor_clone = Arc::clone(&monitor);
    task::spawn(async move {
        monitor_clone.run().await;
    });

    // 保持主线程运行，以便监控器能够持续工作
    loop {
        tokio::time::sleep(Duration::from_secs(60)).await;
    }
}
