package api

import (
	"context"
	"fmt"
	"github.com/atopx/clever"
	"github.com/goccy/go-json"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"wargod/game"
	"wargod/model"
	"wargod/opgg"
)

type Api struct {
	Ctx  context.Context
	Game *game.Game
}

func New() *Api {
	return &Api{}
}

func (a *Api) GameLauncher() error {
	data, err := os.ReadFile(`C:\ProgramData\Riot Games\RiotClientInstalls.json`)
	if err != nil {
		return fmt.Errorf("未寻找到游戏安装路径")
	}
	result := make(map[string]any)
	_ = json.Unmarshal(data, &result)
	clientPath := path.Join(path.Dir(path.Dir(result["rc_default"].(string))), "TCLS", "client.exe")
	output, err := exec.Command(clientPath).CombinedOutput()
	if err != nil {
		return fmt.Errorf(clever.String(output))
	}
	return nil
}

func (a *Api) GetState() string {
	return a.Game.Flow
}

func (a *Api) SetRune(champRune opgg.Rune) {
	// 获取OPGG符文
	data, err := a.Game.Client.Get(game.EndpointPerksCurrentPage)
	if err != nil {
		slog.Error("get current rune page failed", slog.String("err", err.Error()))
		return
	}
	r := model.NewRune(data)
	r.Name = "wargod - atopx"
	r.Order = 0
	r.PrimaryStyleId = champRune.PrimaryPageId
	r.SubStyleId = champRune.SecondaryPageId
	r.Current = true
	r.IsActive = true
	r.SelectedPerkIds = make([]int, 0, 9)
	r.SelectedPerkIds = append(r.SelectedPerkIds, champRune.PrimaryRuneIds...)
	r.SelectedPerkIds = append(r.SelectedPerkIds, champRune.SecondaryRuneIds...)
	r.SelectedPerkIds = append(r.SelectedPerkIds, champRune.StatModIds...)
	_, err = a.Game.Client.Put(fmt.Sprintf(game.EndpointPerksSetPage, r.Id), r.MustMarshal())
	if err != nil {
		slog.Error("update run page failed", slog.String("err", err.Error()))
		return
	}
	slog.Info("设置符文: OK")
}
