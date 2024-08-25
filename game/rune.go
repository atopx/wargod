package game

import (
	"fmt"
	"log/slog"
	"wargod/model"
	"wargod/opgg"
)

func (g *Game) SetRune(champId int) error {

	identity := fmt.Sprintf("atopx - %d", champId)

	// 获取当前符文页
	data, err := g.Client.Get(EndpointPerksCurrentPage)
	if err != nil {
		return fmt.Errorf("get current rune page failed: %w", err)
	}
	r := model.NewRune(data)

	if r.Name == identity {
		slog.Info("英雄未改变, 跳过符文设置")
		return nil
	}

	// 获取OPGG符文
	info := opgg.GetAramChampionInfo(champId)
	// 根据算法获取最佳符文
	opggBestRune := info.Data.BestRune()
	// 设置符文页
	r.Name = identity
	r.Order = 0
	r.PrimaryStyleId = opggBestRune.PrimaryPageId
	r.SubStyleId = opggBestRune.SecondaryPageId
	r.Current = true
	r.IsActive = true
	r.SelectedPerkIds = make([]int, 0, 9)
	r.SelectedPerkIds = append(r.SelectedPerkIds, opggBestRune.PrimaryRuneIds...)
	r.SelectedPerkIds = append(r.SelectedPerkIds, opggBestRune.SecondaryRuneIds...)
	r.SelectedPerkIds = append(r.SelectedPerkIds, opggBestRune.StatModIds...)
	_, err = g.Client.Put(fmt.Sprintf(EndpointPerksSetPage, r.Id), r.MustMarshal())
	if err != nil {
		return fmt.Errorf("update rune page failed: %w", err)
	}
	slog.Info("设置符文", slog.String("rune", identity))
	return nil
}
