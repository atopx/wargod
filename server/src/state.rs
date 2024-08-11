use std::sync::Arc;

use once_cell::sync::Lazy;
use serde::Deserialize;
use serde::Serialize;
use tokio::sync::RwLock;

/// GameState 游戏状态
#[derive(Debug, Serialize, Deserialize, Default, Clone, PartialEq)]
pub enum GameState {
    #[default]
    None,
    // 大厅
    Lobby,
    // 房间
    Room,
    // 匹配中
    Matching,
    // 角色选择
    PreGame,
    // 游戏中
    Gaming,
    // 掉线
    Disconnected,
}

/// GameEvent 游戏事件
pub enum GameEvent {
    // 游戏启动失败 -> None
    LaunchFailed,
    // 游戏启动成功 -> 大厅
    LaunchSuccess,
    // 加入房间 -> 房间
    JoinRoom,
    // 开始匹配 -> 匹配中
    StartMatch,
    // 匹配成功 -> 角色选择
    Matched,
    // 进入游戏 -> 游戏中
    StartGame,
    // 游戏结束 -> 大厅
    EndGame,
    // 断开连接 -> 掉线
    Disconnect,
    // 重新连接 -> 掉线前的状态
    Reconnect,
}

/// GameStateMachine 游戏状态机
pub struct GameStateMachine {
    // 当前状态
    current_state: GameState,
    // 用于状态的逆转，比如离线重连
    previous_state: Option<GameState>,
}

impl GameStateMachine {
    pub fn new() -> Self { Self { current_state: GameState::None, previous_state: None } }

    pub fn transition_to(&mut self, new_state: GameState) {
        println!("Transitioning from {:?} to {:?}", self.current_state, new_state);
        self.previous_state = Some(self.current_state.clone()); // 记录上一个状态
        self.current_state = new_state;
    }

    pub fn handle_event(&mut self, event: GameEvent) {
        match event {
            GameEvent::LaunchFailed => self.transition_to(GameState::None),
            GameEvent::LaunchSuccess => self.transition_to(GameState::Lobby),
            GameEvent::JoinRoom => self.transition_to(GameState::Room),
            GameEvent::StartMatch => self.transition_to(GameState::Matching),
            GameEvent::Matched => self.transition_to(GameState::PreGame),
            GameEvent::StartGame => self.transition_to(GameState::Gaming),
            GameEvent::EndGame => self.transition_to(GameState::Lobby),
            GameEvent::Disconnect => self.transition_to(GameState::Disconnected),
            GameEvent::Reconnect => {
                if let Some(prev_state) = &self.previous_state {
                    self.transition_to(prev_state.clone());
                } else {
                    println!("No previous state to revert to");
                }
            }
        }
    }

    // 获取当前状态
    pub fn get_state(&self) -> &GameState { &self.current_state }
}

// GAME_STATE_MACHINE 定义全局状态机
static GAME_STATE_MACHINE: Lazy<Arc<RwLock<GameStateMachine>>> =
    Lazy::new(|| Arc::new(RwLock::new(GameStateMachine::new())));

#[tokio::test]
async fn test_game_state_machine() {
    {
        // 游戏启动成功 -> 大厅
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::LaunchSuccess);
        assert_eq!(state_machine.get_state(), &GameState::Lobby);
    } // 写锁在这里释放

    {
        // 加入房间 -> 房间
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::JoinRoom);
        assert_eq!(state_machine.get_state(), &GameState::Room);
    }

    {
        // 开始匹配 -> 匹配中
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::StartMatch);
        assert_eq!(state_machine.get_state(), &GameState::Matching);
    }

    {
        // 匹配成功 -> 角色选择
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::Matched);
        assert_eq!(state_machine.get_state(), &GameState::PreGame);
    }

    {
        // 进入游戏 -> 游戏中
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::StartGame);
        assert_eq!(state_machine.get_state(), &GameState::Gaming);
    }

    {
        // 断开连接 -> 掉线
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::Disconnect);
        assert_eq!(state_machine.get_state(), &GameState::Disconnected);
    }

    {
        // 重新连接 -> 游戏中
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::Reconnect);
        assert_eq!(state_machine.get_state(), &GameState::Gaming);
    }

    {
        // 游戏结束 -> 大厅
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::EndGame);
        assert_eq!(state_machine.get_state(), &GameState::Lobby);
    }

    {
        // 测试从断开连接状态重连回原来的状态
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::JoinRoom); // 先加入房间
        state_machine.handle_event(GameEvent::Disconnect); // 然后断开连接
        assert_eq!(state_machine.get_state(), &GameState::Disconnected);

        state_machine.handle_event(GameEvent::Reconnect); // 重新连接后应回到房间状态
        assert_eq!(state_machine.get_state(), &GameState::Room);
    }

    {
        // 游戏启动失败 -> None
        let mut state_machine = GAME_STATE_MACHINE.write().await;
        state_machine.handle_event(GameEvent::LaunchFailed);
        assert_eq!(state_machine.get_state(), &GameState::None);
    }
}
