use std::error::Error;
use std::fmt;
use std::fmt::Display;

/// Errors that can occur when trying to get the Riot process information
#[derive(Debug, Clone)]
pub enum ProcessError {
    /// League client has not been started
    ProcessNotAvailable,
    /// There has been an error read the process info
    ProcessAccessDenied,
}

impl Error for ProcessError {}

impl Display for ProcessError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Self::ProcessNotAvailable => {
                write!(f, "{:?}: Riot/League client process could not be found", self)
            }
            Self::ProcessAccessDenied => {
                write!(f, "{:?}: Read League client process denied", self)
            }
        }
    }
}

/// Errors for the Ingame API
#[derive(Debug, Clone)]
pub enum WatchError {
    /// An API might not be available yet during the loading screen
    ApiNotAvailableInSpectatorMode,
    ApiNotAvailableDuringLoadingScreen,
    /// An error occurred on the client side probably because of a malformed request \
    /// Corresponds to HTTP status responses 400 – 499, excluding 400 and 404 which are [WatchError::ApiNotAvailableInSpectatorMode] and [WatchError::ApiNotAvailableDuringLoadingScreen]
    ClientError(String),
    /// An error ocurred on the server side \
    /// Corresponds to HTTP status responses 500 – 599
    ServerError(String),
    /// There was an error deserializing the received data
    DeserializationError(String),
    /// All errors not caught by the other [WatchError] variants are categorised as a [WatchError::ConnectionError]
    ConnectionError(String),
}

impl From<reqwest::Error> for WatchError {
    fn from(error: reqwest::Error) -> Self {
        if let Some(status) = error.status() {
            if status == 400 {
                return WatchError::ApiNotAvailableInSpectatorMode;
            } else if status == 404 {
                return WatchError::ApiNotAvailableDuringLoadingScreen;
            } else if status.is_client_error() {
                return WatchError::ClientError(status.to_string());
            } else if status.is_server_error() {
                return WatchError::ServerError(status.to_string());
            }
        }
        if error.is_decode() {
            return WatchError::DeserializationError(error.to_string());
        }
        WatchError::ConnectionError(error.to_string())
    }
}

impl Error for WatchError {}
impl Display for WatchError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result { fmt::Debug::fmt(self, f) }
}
