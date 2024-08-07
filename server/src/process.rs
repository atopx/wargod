use std::ffi::OsString;
use std::mem::size_of;
use std::mem::zeroed;
use std::os::windows::prelude::OsStringExt;
use std::ptr::null_mut;

use windows::core::PCWSTR;
use windows::core::PWSTR;
use windows::Wdk::System::Threading::NtQueryInformationProcess;
use windows::Wdk::System::Threading::PROCESSINFOCLASS;
use windows::Win32::Foundation::CloseHandle;
use windows::Win32::Foundation::GetLastError;
use windows::Win32::Foundation::BOOL;
use windows::Win32::Foundation::HANDLE;
use windows::Win32::Foundation::HWND;
use windows::Win32::Security::GetTokenInformation;
use windows::Win32::Security::TokenElevation;
use windows::Win32::Security::TOKEN_ELEVATION;
use windows::Win32::Security::TOKEN_QUERY;
use windows::Win32::System::Diagnostics::Debug::ReadProcessMemory;
use windows::Win32::System::Diagnostics::ToolHelp::CreateToolhelp32Snapshot;
use windows::Win32::System::Diagnostics::ToolHelp::Process32FirstW;
use windows::Win32::System::Diagnostics::ToolHelp::Process32NextW;
use windows::Win32::System::Diagnostics::ToolHelp::PROCESSENTRY32W;
use windows::Win32::System::Diagnostics::ToolHelp::TH32CS_SNAPPROCESS;
use windows::Win32::System::Threading::GetCurrentProcess;
use windows::Win32::System::Threading::OpenProcess;
use windows::Win32::System::Threading::OpenProcessToken;
use windows::Win32::System::Threading::PEB;
use windows::Win32::System::Threading::PROCESS_BASIC_INFORMATION;
use windows::Win32::System::Threading::PROCESS_QUERY_LIMITED_INFORMATION;
use windows::Win32::System::Threading::PROCESS_VM_READ;
use windows::Win32::System::Threading::RTL_USER_PROCESS_PARAMETERS;
use windows::Win32::UI::Shell::ShellExecuteW;
use windows::Win32::UI::WindowsAndMessaging::SW_HIDE;

use super::error::ProcessError;

/// return the first process id by the name you gave
pub fn find_process_id_by_name(process_name: &str) -> Result<u32, ProcessError> {
    unsafe {
        let handle = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0).unwrap();
        let mut process = zeroed::<PROCESSENTRY32W>();
        process.dwSize = size_of::<PROCESSENTRY32W>() as u32;
        let end = process_name.len();

        let result = match Process32FirstW(handle, &mut process) {
            Ok(_) => loop {
                match Process32NextW(handle, &mut process) {
                    Ok(_) => {
                        let find_name = String::from_utf16_lossy(&process.szExeFile[0..end]);
                        if process_name == find_name {
                            break Ok(process.th32ProcessID);
                        }
                    }
                    Err(_) => break Err(ProcessError::ProcessNotAvailable),
                }
            },
            Err(_) => Err(ProcessError::ProcessNotAvailable),
        };

        let _ = CloseHandle(handle);
        result
    }
}

/// get the process command line params
pub fn get_process_cmdline(pid: u32) -> Result<String, ProcessError> {
    unsafe {
        match OpenProcess(PROCESS_QUERY_LIMITED_INFORMATION | PROCESS_VM_READ, BOOL(0), pid) {
            Ok(h) => {
                let pc = zeroed::<PROCESSINFOCLASS>();
                let mut pbi = zeroed::<PROCESS_BASIC_INFORMATION>();
                if NtQueryInformationProcess(
                    h,
                    pc,
                    std::ptr::addr_of_mut!(pbi) as _,
                    size_of::<PROCESS_BASIC_INFORMATION>() as _,
                    null_mut(),
                )
                .is_ok()
                {
                    let mut peb = zeroed::<PEB>();
                    match ReadProcessMemory(
                        h,
                        pbi.PebBaseAddress as _,
                        std::ptr::addr_of_mut!(peb) as _,
                        size_of::<PEB>(),
                        Some(null_mut()),
                    ) {
                        Ok(_) => {
                            let mut proc_params = zeroed::<RTL_USER_PROCESS_PARAMETERS>();

                            match ReadProcessMemory(
                                h,
                                peb.ProcessParameters as _,
                                std::ptr::addr_of_mut!(proc_params) as _,
                                size_of::<RTL_USER_PROCESS_PARAMETERS>(),
                                Some(null_mut()),
                            ) {
                                Ok(_) => {
                                    let cmd_lenth = proc_params.CommandLine.MaximumLength;
                                    let cmd_buffer = proc_params.CommandLine.Buffer;

                                    Ok(read_process_buffer(cmd_buffer, cmd_lenth, h))
                                }
                                Err(_) => {
                                    let _ = CloseHandle(h);
                                    Err(ProcessError::ProcessAccessDenied)
                                }
                            }
                        }
                        Err(_) => {
                            let _ = CloseHandle(h);
                            Err(ProcessError::ProcessAccessDenied)
                        }
                    }
                } else {
                    Err(ProcessError::ProcessAccessDenied)
                }
            }
            Err(_) => Err(ProcessError::ProcessAccessDenied),
        }
    }
}

/// this function is used to get transfer `PWSTR` of the process params to `String` . it need to do `ReadProcessMemory` again.
unsafe fn read_process_buffer(pwstr: PWSTR, len: u16, h: HANDLE) -> String {
    // Create a buffer with the same length
    let mut temp = vec![0u16; len as usize];

    // Read process memory into buffer
    let success = ReadProcessMemory(
        h,
        pwstr.0 as _,
        temp.as_mut_ptr() as _,
        (len as usize * std::mem::size_of::<u16>()) as _,
        Some(null_mut()),
    );

    if success.is_ok() {
        // Find the null terminator in the buffer
        let x = &temp[..len as usize];

        let cmd_params = match x.iter().position(|&c| c == 0) {
            Some(nul) => OsString::from_wide(&x[..nul]),
            None => OsString::from_wide(x),
        }
        .into_string()
        .unwrap_or_else(|_| "Failed to convert to String".to_string());
        cmd_params
    } else {
        let _ = CloseHandle(h);
        format!("access denied {:?}", GetLastError())
    }
}

/// detect if it was run by an administrator
pub fn is_admin() -> bool {
    unsafe {
        let mut token_handle: HANDLE = HANDLE(null_mut());
        let process_handle = GetCurrentProcess();
        if OpenProcessToken(process_handle, TOKEN_QUERY, &mut token_handle).is_ok() {
            let mut elevation = TOKEN_ELEVATION { TokenIsElevated: 0 };
            let mut return_length: u32 = 0;
            let success = GetTokenInformation(
                token_handle,
                TokenElevation,
                Option::from(&mut elevation as *mut _ as *mut core::ffi::c_void),
                size_of::<TOKEN_ELEVATION>() as u32,
                &mut return_length,
            )
            .is_ok();
            let _ = CloseHandle(token_handle);
            if success {
                return elevation.TokenIsElevated != 0;
            }
        }
        false
    }
}

/// The request is re-run with administrator privileges
pub fn elevated() {
    let exe = std::env::current_exe().unwrap();
    let exe_str = exe.to_str().unwrap();
    let cwd = std::env::current_dir().unwrap();
    let cwd_str = cwd.to_str().unwrap();
    let args: String = std::env::args().skip(1).collect::<Vec<String>>().join(" ");
    let exe_wide: Vec<u16> = exe_str.encode_utf16().chain(Some(0)).collect();
    let cwd_wide: Vec<u16> = cwd_str.encode_utf16().chain(Some(0)).collect();
    let args_wide: Vec<u16> = args.encode_utf16().chain(Some(0)).collect();
    let runas_wide: Vec<u16> = "runas".encode_utf16().chain(Some(0)).collect();

    unsafe {
        let result = ShellExecuteW(
            HWND(null_mut()),
            PCWSTR(runas_wide.as_ptr()),
            PCWSTR(exe_wide.as_ptr()),
            PCWSTR(args_wide.as_ptr()),
            PCWSTR(cwd_wide.as_ptr()),
            SW_HIDE,
        );
        if result.0 as isize <= 32 {
            println!("Unable to run program with administrator privileges, please check UAC settings");
            std::process::exit(1);
        }
    }

    std::process::exit(0);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_flow() {
        assert!(!is_admin());
        elevated();
        assert!(is_admin());

        let pid_result = find_process_id_by_name("LeagueClientUx.exe");
        assert!(pid_result.is_ok());

        let cmdline_result = get_process_cmdline(pid_result.unwrap());
        assert!(cmdline_result.is_ok());

        assert!(cmdline_result.unwrap().contains("--remoting-auth-token"));
    }
}
