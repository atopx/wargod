// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

pub mod error;
pub mod process;

fn init_logging() { env_logger::Builder::from_default_env().filter_level(log::LevelFilter::Info).init(); }

fn main() {
    init_logging();
    wargod_lib::run()
}
