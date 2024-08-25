package game

import "github.com/atopx/clever"

const (
	EndpointGameFlowPhase              = "/lol-gameflow/v1/gameflow-phase"        // 游戏状态机
	EndpointChampSelectSession         = "/lol-champ-select/v1/session"           // 英雄选择动态
	EndpointMatchReadyAccept           = "/lol-matchmaking/v1/ready-check/accept" // 接受对局
	EndpointChampSelectCurrentChampion = "/lol-champ-select/v1/current-champion"  // 当前选择英雄
	EndpointChatMe                     = "/lol-chat/v1/me"                        // 当前召唤师
	EndpointPerksCurrentPage           = "/lol-perks/v1/currentpage"
	EndpointPerksSetPage               = "/lol-perks/v1/pages/%d"
)

var (
	EmptyJsonBody = clever.Bytes(`{}`)
)

const (
	FlowNone            = "None"            // 大厅无状态
	FlowLobby           = "Lobby"           // 进入房间
	FlowMatchmaking     = "Matchmaking"     // 开始匹配
	FlowReadyCheck      = "ReadyCheck"      // 匹配成功
	FlowChampSelect     = "ChampSelect"     // 英雄选择
	FlowGameStart       = "GameStart"       // 游戏开始
	FlowInProgress      = "InProgress"      // 载入游戏
	FlowWaitingForStats = "WaitingForStats" // 等待游戏结果判定
	FlowPreEndOfGame    = "PreEndOfGame"    // 游戏结束
	FlowEndOfGame       = "EndOfGame"       // 退出游戏
	Disconnect          = "Disconnect"      // 掉线
)
