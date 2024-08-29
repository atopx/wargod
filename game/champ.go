package game

import (
	"fmt"
	"github.com/atopx/clever"
	"github.com/goccy/go-json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log/slog"
	"wargod/conf"
	"wargod/model"
	"wargod/opgg"
	"wargod/types"
)

// SwapChampion 发送备战席交换请求
func (g *Game) SwapChampion(champId int) error {
	swapResp, code, err := g.Client.Post(fmt.Sprintf("/lol-champ-select/v1/session/bench/swap/%d", champId), nil)
	if err != nil {
		return fmt.Errorf("swap champ %d: %w", champId, err)
	}
	if code >= 400 {
		return fmt.Errorf("swap champ %d: %s", champId, clever.String(swapResp))
	}
	return nil
}

// UseDices 摇骰子, num为0时使用所有骰子
func (g *Game) UseDices(num int) error {
	for i := 0; i < num; i++ {
		_, _, err := g.Client.Post("/lol-champ-select/v1/session/my-selection/reroll", EmptyJsonBody)
		if err != nil {
			return fmt.Errorf("使用骰子失败: %w", err)
		}
	}
	return nil
}

// GetSwapChampionId 匹配备战席中最高优先级的英雄
func (g *Game) GetSwapChampionId(champ model.Team, benchChampions []model.BenchChampion) int {
	truncation := len(conf.Entry.AramConfig.SwapPriorityChampIds)
	// 1. 判断当前英雄在优先级队列的位置 currentIdx, 如果不存在取SwapPriorityChampIds长度
	for i, priority := range conf.Entry.AramConfig.SwapPriorityChampIds {
		if priority == champ.ChampionId {
			truncation = i
			break
		}
	}
	// 2. 如果当前英雄不是最高优先级, 则进行逻辑处理
	if truncation > 0 {
		// 3. 获取最高优先级的英雄ID, 如果存在则交换:
		// - 外层遍历本地优先级队列 { 内层遍历游戏备战席 }
		// - 第一个匹配到的则为最高优先级
		for _, priority := range conf.Entry.AramConfig.SwapPriorityChampIds[:truncation] {
			for _, bench := range benchChampions {
				if bench.ChampionId == priority {
					return bench.ChampionId
				}
			}
		}

	}
	return -1
}

func (g *Game) champSelectHandle(data []byte) error {
	var resp model.EventMessageBlob[model.ChampSelect]

	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("parse champ select event failed: %w", err)
	}

	if resp.EventType == "Delete" {
		return nil
	}

	champ := resp.Data.MyTeam[resp.Data.LocalPlayerCellId%5]

	// 启用大乱斗自动交换
	if conf.Entry.AutoSwap &&
		len(conf.Entry.AramConfig.SwapPriorityChampIds) > 0 &&
		resp.Data.BenchEnabled &&
		len(resp.Data.BenchChampions) > 0 {

		// 先使用所有骰子
		if err := g.UseDices(1); err != nil {
			slog.Error(err.Error())
		}

		// 寻找高优先级英雄
		if champId := g.GetSwapChampionId(champ, resp.Data.BenchChampions); champId > 0 {
			// 这里可以直接结束函数, 切换后会触发下一个事件
			if c, ok := model.ChampMap[champId]; ok {
				types.NewNotice(g.ctx, "success", fmt.Sprintf("自动切换: [%s - %s]", c.Title, c.Name))
			}
			return g.SwapChampion(champId)
		}
	}

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
