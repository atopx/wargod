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
	"wargod/conf"
	"wargod/game"
	"wargod/model"
	"wargod/opgg"
	"wargod/types"
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

func (a *Api) GetConfig() *conf.Config {
	return conf.Entry
}

func (a *Api) SetConfig(entry *conf.Config) error {
	conf.Entry.AutoNext = entry.AutoNext
	conf.Entry.AutoAccept = entry.AutoAccept
	conf.Entry.AutoStatus = entry.AutoStatus
	conf.Entry.AutoSwap = entry.AutoSwap
	conf.Entry.StatusContent = entry.StatusContent
	conf.Entry.AramConfig = entry.AramConfig
	if conf.Entry.StatusContent.RankedLeagueTier == types.TierNone ||
		conf.Entry.StatusContent.RankedLeagueTier == types.TierMaster ||
		conf.Entry.StatusContent.RankedLeagueTier == types.TierGrandmaster ||
		conf.Entry.StatusContent.RankedLeagueTier == types.TierChallenger {
		conf.Entry.StatusContent.RankedLeagueDivision = types.TierDivision0
	}
	if conf.Entry.AutoStatus {
		_, _, err := a.Game.Client.Put(game.EndpointChatMe, conf.Entry.StatusContent.Data())
		if err != nil {
			slog.Error("set chat-me status failed", slog.String("err", err.Error()))
		}
	}

	if conf.Entry.AutoNext && a.Game.Flow == game.FlowLobby {
		// 开启自动续盘且在房间状态中, 自动开始
		if err := a.Game.StartLobbyMatchmaking(); err != nil {
			slog.Error("start lobby matchmaking failed", slog.String("err", err.Error()))
		}
	}
	return conf.SaveConfig()
}

func (a *Api) TierMeta() ([]types.RankedQueue, []types.Tier, []types.TierDivision) {
	fmt.Println(types.AllQueues)
	return types.AllQueues, types.AllTiers, types.AllTierDivision
}

func (a *Api) GetState() string {
	return a.Game.Flow
}

func (a *Api) SetRune(champRune opgg.Rune) {
	// 获取OPGG符文
	data, _, err := a.Game.Client.Get(game.EndpointPerksCurrentPage)
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
	_, _, err = a.Game.Client.Put(fmt.Sprintf(game.EndpointPerksSetPage, r.Id), r.MustMarshal())
	if err != nil {
		slog.Error("update run page failed", slog.String("err", err.Error()))
		return
	}
	slog.Info("设置符文: OK")
}

// CurrentSummoner 查询召唤师信息
func (a *Api) CurrentSummoner() (*model.Summoner, error) {
	data, _, err := a.Game.Client.Get("/lol-summoner/v1/current-summoner")
	if err != nil {
		return nil, fmt.Errorf("get current summoner failed: %w", err)
	}
	summoner := new(model.Summoner)
	_ = json.Unmarshal(data, summoner)
	return summoner, nil
}

// SearchSummoner 搜索召唤师信息
func (a *Api) SearchSummoner(name string) (*model.Summoner, error) {
	data, code, err := a.Game.Client.Get("/lol-summoner/v1/summoners?name=" + name)
	if err != nil {
		return nil, fmt.Errorf("get summoner failed: %w", err)
	}
	if code == 422 {
		return nil, fmt.Errorf("未搜索到召唤师")
	}
	summoner := new(model.Summoner)
	_ = json.Unmarshal(data, summoner)
	return summoner, nil
}

// GetSummoner 获取指定召唤师信息
func (a *Api) GetSummoner(puuid string) (*model.Summoner, error) {
	data, code, err := a.Game.Client.Get("/lol-summoner/v2/summoners/puuid/" + puuid)
	if err != nil {
		return nil, fmt.Errorf("get summoner failed: %w", err)
	}
	if code >= 400 {
		return nil, fmt.Errorf("无效的PUUID: %s", puuid)
	}
	summoner := new(model.Summoner)
	_ = json.Unmarshal(data, summoner)
	return summoner, nil
}

// ListGameSummary 游戏战绩列表
func (a *Api) ListGameSummary(puuid string, start, end int) (*model.GameSummary, error) {
	data, _, err := a.Game.Client.Get(fmt.Sprintf("/lol-match-history/v1/products/lol/%s/matches?begIndex=%d&endIndex=%d", puuid, start, end))
	if err != nil {
		return nil, fmt.Errorf("get summoner game list failed: %w", err)
	}
	record := new(model.GameSummary)
	_ = json.Unmarshal(data, &record)
	return record, nil
}

// GetGameInfo 对局详情
func (a *Api) GetGameInfo(gameId int64) (*model.GameInfo, error) {
	data, code, err := a.Game.Client.Get(fmt.Sprintf("/lol-match-history/v1/games/%d", gameId))
	if err != nil {
		return nil, fmt.Errorf("get game info failed: %w", err)
	}
	if code >= 400 {
		return nil, fmt.Errorf("无效的GameId: %d", gameId)
	}
	record := new(model.GameInfo)
	_ = json.Unmarshal(data, &record)
	return record, nil
}

// GetRankInfo 召唤师段位数据
func (a *Api) GetRankInfo(puuid string) (*model.Ranked, error) {
	data, code, err := a.Game.Client.Get("/lol-ranked/v1/ranked-stats/" + puuid)
	if err != nil {
		return nil, fmt.Errorf("get ranking info failed: %w", err)
	}
	if code >= 400 {
		return nil, fmt.Errorf("无效的PUUID: %s", puuid)
	}
	ranked := new(model.Ranked)
	_ = json.Unmarshal(data, &ranked)
	return ranked, nil
}
