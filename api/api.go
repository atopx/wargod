package api

import (
	"context"
	"fmt"
	"log/slog"
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
