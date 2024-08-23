package main

import (
	"context"
	"fmt"
	"github.com/atopx/clever"
	"github.com/goccy/go-json"
	"log"
	"log/slog"
	"time"
	"wargod/core"
	"wargod/model"
)

type EventType string

var (
	EventTypeUpdate = "Update"
	EventTypeCreate = "Create"
	EventTypeDelete = "Delete"
)

type EventMessageBlob[T any] struct {
	EventType string `json:"eventType"`
	Uri       string `json:"uri"`
	Data      T      `json:"data"`
}

const (
	GameFlowUri     = "/lol-gameflow/v1/gameflow-phase"
	ChampSessionUri = "/lol-champ-select/v1/session"

	defaultSummonerStatus = `{"statusMessage":"Garbage games","lol":{"rankedLeagueQueue":"RANKED_SOLO_5x5","rankedLeagueTier":"CHALLENGER"}}`
)

var client *core.Client

func GameFlowHandle(data []byte) error {
	var resp EventMessageBlob[string]
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("parse game flow event failed: %w", err)
	}
	switch resp.Data {
	case "None":
		slog.Info("大厅无状态")
	case "Lobby":
		slog.Info("进入房间")
	case "Matchmaking":
		slog.Info("开始匹配")
	case "ChampSelect":
		// 获取选择的英雄
		slog.Info("获取英雄")
	case "ReadyCheck":
		slog.Info("匹配成功")
		if _, err := client.Post("/lol-matchmaking/v1/ready-check/accept", []byte(`{}`)); err != nil {
			slog.Error(err.Error())
		}
	case "GameStart":
		// 确定英雄, 推荐出装
		respData, _ := client.Get("/lol-champ-select/v1/current-champion")
		slog.Info("锁定英雄: " + string(respData))
	case "InProgress":
		slog.Info("载入游戏中...")

	case "WaitingForStats":
		slog.Info("结束, 等待游戏结果判定")
	case "PreEndOfGame":
		slog.Info("等待确认后推出游戏")
	case "EndOfGame":
		slog.Info("游戏结束")
		// 异步分析战绩
	default:
		slog.Info("忽略状态: {}" + resp.Data)
	}
	return nil
}

func ChampSelectHandle(data []byte) error {
	var resp EventMessageBlob[model.ChampSelect]

	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("parse champ select event failed: %w", err)
	}

	if resp.EventType == "Delete" {
		return nil
	}

	champ := resp.Data.MyTeam[resp.Data.LocalPlayerCellId%5]
	record := model.ChampMap[champ.ChampionId]
	opgg := fmt.Sprintf("https://www.op.gg/modes/aram/%s/build", record.Alias)
	slog.Info("select champ", slog.String("title", record.Title), slog.String("Link", opgg), slog.String("resp.EventType", resp.EventType))
	// TODO 自动配置符文
	return nil
}

func main() {
	p, err := core.NewSearch().WaitForLauncher()
	if err != nil {
		log.Fatal(err)
	}
	ga, err := core.NewGameAuth(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", ga)
	ctx := context.Background()
	for {
		client, err = core.New(ctx, time.Second, ga)
		if err == nil {
			break
		}
		slog.Info(err.Error())
		time.Sleep(time.Second)
	}

	if err = client.Subscribe(map[string]func(data []byte) error{
		GameFlowUri:     GameFlowHandle,
		ChampSessionUri: ChampSelectHandle,
	}); err != nil {
		log.Fatal(err)
	}
	// 设置状态
	_, err = client.Put("/lol-chat/v1/me", clever.Bytes(defaultSummonerStatus))
	if err != nil {
		log.Fatal(err)
	}
	for {
		client.Listen()
		// 如果websocket出问题, 大概率是游戏掉线, 重新监听
		main()
	}
}
