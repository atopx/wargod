package game

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"
	"wargod/conf"
	"wargod/core"
	"wargod/model"

	"github.com/atopx/clever"
	"github.com/goccy/go-json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Game struct {
	ctx    context.Context
	auth   *core.GameAuth
	Client *core.Client
	Flow   string
	Mode   string
	champ  int
}

func New(ctx context.Context) *Game {
	return &Game{ctx: ctx, Flow: Disconnect}
}

func (g *Game) start() error {
	p, err := core.NewSearch().WaitForLauncher()
	if err != nil {
		return fmt.Errorf("wait for game launcher error: %w", err)
	}

	g.auth, err = core.NewGameAuth(p)
	if err != nil {
		return fmt.Errorf("get game auth error: %w", err)
	}
	g.waitForConnect()

	if err = g.setDefaultStatus(); err != nil {
		return fmt.Errorf("set game default status error: %w", err)
	}

	return g.subscribe()
}

func (g *Game) Start() {
	for {
		for {
			err := g.start()
			if err == nil {
				break
			}
			slog.Error(err.Error())
			time.Sleep(1 * time.Second)
		}
		runtime.EventsEmit(g.ctx, "GameState", FlowNone)
		g.Client.Listen()
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) waitForConnect() {
	var err error
	for {
		g.Client, err = core.New(g.ctx, time.Second, g.auth)
		if err == nil {
			break
		}
		slog.Info("connect to game server error", slog.String("error", err.Error()))
		time.Sleep(time.Second)
	}
}

func (g *Game) subscribe() error {
	// 订阅游戏事件
	return g.Client.Subscribe(map[string]func(data []byte) error{
		EndpointGameFlowPhase:      g.gameFlowHandle,
		EndpointChampSelectSession: g.champSelectHandle,
	})
}

func (g *Game) setDefaultStatus() error {
	// 进入游戏后的一些默认设置, 比如默认设置为最强王者 hahaha
	if conf.Entry.AutoStatus {
		_, _, err := g.Client.Put(EndpointChatMe, conf.Entry.StatusContent.Data())
		if err != nil {
			return err
		}
	}
	data, _, _ := g.Client.Get(EndpointGameFlowPhase)
	g.Flow = strings.Trim(clever.String(data), `"`)
	runtime.WindowSetTitle(g.ctx, "极地战神 - "+FlowCnMap[g.Flow])
	return nil
}

func (g *Game) GetGameMode() {
	data, _, err := g.Client.Get(EndpointLobby)
	if err != nil {
		slog.Error("get game lobby error", slog.String("error", err.Error()))
		return
	}
	var lobby model.Lobby
	_ = json.Unmarshal(data, &lobby)
	g.Mode = lobby.GameConfig.GameMode
}
