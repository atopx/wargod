package game

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log/slog"
	"wargod/model"
)

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
		//g.GetGameMode()
		slog.Info("进入房间")
	case FlowMatchmaking:
		slog.Info("开始匹配")
	case FlowChampSelect:
		// 获取选择的英雄
		slog.Info("获取英雄")
	case FlowReadyCheck:
		slog.Info("匹配成功")
		if _, err := g.Client.Post(EndpointMatchReadyAccept, EmptyJsonBody); err != nil {
			slog.Error(err.Error())
		}
	case FlowGameStart:
		// 确定英雄, 推荐出装
		respData, _ := g.Client.Get(EndpointChampSelectCurrentChampion)
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
