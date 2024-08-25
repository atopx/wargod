package game

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log/slog"
	"wargod/model"
	"wargod/opgg"
)

// 发送英雄选择事件

func (g *Game) champSelectHandle(data []byte) error {
	var resp model.EventMessageBlob[model.ChampSelect]

	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("parse champ select event failed: %w", err)
	}

	if resp.EventType == "Delete" {
		return nil
	}

	champ := resp.Data.MyTeam[resp.Data.LocalPlayerCellId%5]

	if g.champ != champ.ChampionId {
		// 如果英雄发生变化发送事件到前端
		record := model.ChampMap[champ.ChampionId]
		link := fmt.Sprintf("https://www.op.gg/modes/aram/%s/build", record.Alias)
		slog.Info("select champ", slog.String("name", record.Name), slog.String("Link", link), slog.String("resp.EventType", resp.EventType))
		// TODO 这里先写死大乱斗模式
		opggInfo := opgg.GetAramChampionInfo(champ.ChampionId)

		skillMaster := opggInfo.Data.SkillMasteries[0]
		skillMasterOrder := opggInfo.Data.SkillMasteries[0].Builds[0]

		// 转换为Message
		message := ChampChangeMessage{
			Champion: ChampChangeMessageChampion{
				Id:       champ.ChampionId,
				Title:    record.Title,
				Name:     record.Name,
				Roles:    record.Roles,
				WinRate:  opggInfo.Data.Summary.AverageStats.WinRate,
				PickRate: opggInfo.Data.Summary.AverageStats.WinRate,
				BanRate:  opggInfo.Data.Summary.AverageStats.WinRate,
				Tier:     opggInfo.Data.Summary.AverageStats.Tier,
				Rank:     opggInfo.Data.Summary.AverageStats.Rank,
				Kda:      opggInfo.Data.Summary.AverageStats.Kda,
			},
			Skill: ChampChangeMessageSkill{
				Ids:      skillMaster.Ids,
				Play:     skillMaster.Play,
				Win:      skillMaster.Win,
				PickRate: skillMasterOrder.PickRate,
				Order:    skillMasterOrder.Order,
			},
		}

		if len(opggInfo.Data.Runes) > 3 {
			message.Runes = opggInfo.Data.Runes[:3]
		} else {
			message.Runes = opggInfo.Data.Runes
		}

		if len(opggInfo.Data.SummonerSpells) > 3 {
			message.Spells = opggInfo.Data.SummonerSpells[:3]
		} else {
			message.Spells = opggInfo.Data.SummonerSpells
		}

		if len(opggInfo.Data.Boots) > 3 {
			message.Boots = opggInfo.Data.Boots[:3]
		} else {
			message.Boots = opggInfo.Data.Boots
		}

		if len(opggInfo.Data.StarterItems) > 3 {
			message.StarterItems = opggInfo.Data.StarterItems[:3]
		} else {
			message.StarterItems = opggInfo.Data.StarterItems
		}

		if len(opggInfo.Data.CoreItems) > 5 {
			message.CoreItems = opggInfo.Data.CoreItems[:5]
		} else {
			message.CoreItems = opggInfo.Data.CoreItems
		}
		runtime.EventsEmit(g.ctx, "ChampSelect", message)
	}
	return nil
}

type ChampChangeMessageChampion struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Name     string   `json:"name"`
	Roles    []string `json:"roles"`
	WinRate  float64  `json:"win_rate"`
	PickRate float64  `json:"pick_rate"`
	BanRate  float64  `json:"ban_rate,omitempty"`
	Tier     int      `json:"tier,omitempty"`
	Rank     int      `json:"rank,omitempty"`
	Kda      float64  `json:"kda"`
}

type ChampChangeMessageSkill struct {
	Ids      []string `json:"ids"`
	Play     int      `json:"play"`
	Win      int      `json:"win"`
	PickRate float64  `json:"pick_rate"`
	Order    []string `json:"order"`
}

type ChampChangeMessage struct {
	Champion     ChampChangeMessageChampion `json:"champion"`
	Skill        ChampChangeMessageSkill    `json:"skill"`
	Runes        []*opgg.Rune               `json:"runes"`
	Spells       []*opgg.ChampionItem       `json:"spells"`
	CoreItems    []*opgg.ChampionItem       `json:"core_items"`
	Boots        []*opgg.ChampionItem       `json:"boots"`
	StarterItems []*opgg.ChampionItem       `json:"starter_items"`
}
