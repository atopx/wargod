package game

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log/slog"
	"wargod/conf"
	"wargod/model"
)

func (g *Game) StartLobbyMatchmaking() error {
	if _, _, err := g.Client.Post("/lol-lobby/v2/lobby/matchmaking/search", EmptyJsonBody); err != nil {
		return fmt.Errorf("自动续盘失败: %w", err)
	}
	return nil
}

func (g *Game) gameFlowHandle(data []byte) error {
	var resp model.EventMessageBlob[string]
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("parse game flow event failed: %w", err)
	}
	// 发送状态事件
	runtime.EventsEmit(g.ctx, "GameState", resp.Data)
	runtime.WindowSetTitle(g.ctx, "极地战神 - "+FlowCnMap[resp.Data])
	g.Flow = resp.Data
	switch resp.Data {
	case FlowNone:
		slog.Info("大厅无状态")
	case FlowLobby:
		// 获取游戏模式
		slog.Info("进入房间", slog.Bool("AutoNext", conf.Entry.AutoNext))
		if conf.Entry.AutoNext {
			if err := g.StartLobbyMatchmaking(); err != nil {
				slog.Error(err.Error())
			}
		}
	case FlowMatchmaking:
		slog.Info("开始匹配")
	case FlowChampSelect:
		// 获取选择的英雄
		slog.Info("获取英雄")
	case FlowReadyCheck:
		slog.Info("匹配成功")
		if conf.Entry.AutoAccept {
			// 如果开启了自动应战
			if _, _, err := g.Client.Post(EndpointMatchReadyAccept, EmptyJsonBody); err != nil {
				slog.Error(err.Error())
			}
		}
	case FlowGameStart:
		// 确定英雄, 推荐出装
		respData, _, _ := g.Client.Get(EndpointChampSelectCurrentChampion)
		slog.Info("锁定英雄: " + string(respData))
	case FlowInProgress:
		slog.Info("载入游戏中...")
	case FlowWaitingForStats:
		slog.Info("结束, 等待游戏结果判定")
	case FlowPreEndOfGame:
		slog.Info("等待确认后退出游戏")
	case FlowEndOfGame:
		slog.Info("游戏结束")
		// 异步分析战绩
	default:
		slog.Info("忽略状态: {}" + resp.Data)
	}
	return nil
}
