package opgg

import (
	"time"
)

type Meta struct {
	Version    string    `json:"version,omitempty"`
	CachedAt   string    `json:"cached_at,omitempty"`
	MatchCount int       `json:"match_count,omitempty"`
	AnalyzedAt time.Time `json:"analyzed_at,omitempty"`
}

type Role struct {
	Name  string `json:"name,omitempty"`
	Stats struct {
		WinRate  float64 `json:"win_rate,omitempty"`
		RoleRate float64 `json:"role_rate,omitempty"`
		Play     int     `json:"play,omitempty"`
		Win      int     `json:"win,omitempty"`
	} `json:"stats,omitempty"`
}

type Position struct {
	Name  string `json:"name,omitempty"`
	Stats struct {
		WinRate  float64 `json:"win_rate,omitempty"`
		PickRate float64 `json:"pick_rate,omitempty"`
		RoleRate float64 `json:"role_rate,omitempty"`
		BanRate  float64 `json:"ban_rate,omitempty"`
		Kda      float64 `json:"kda,omitempty"`
		TierData struct {
			Tier          int `json:"tier,omitempty"`
			Rank          int `json:"rank,omitempty"`
			RankPrev      int `json:"rank_prev,omitempty"`
			RankPrevPatch int `json:"rank_prev_patch,omitempty"`
		} `json:"tier_data,omitempty"`
	} `json:"stats,omitempty"`
	Roles    []Role `json:"roles,omitempty"`
	Counters []struct {
		ChampionId int `json:"champion_id,omitempty"`
		Play       int `json:"play,omitempty"`
		Win        int `json:"win,omitempty"`
	} `json:"counters,omitempty"`
}

type ChampionSummary struct {
	Id           int  `json:"id,omitempty"`
	IsRotation   bool `json:"is_rotation,omitempty"`
	IsRip        bool `json:"is_rip,omitempty"`
	AverageStats struct {
		WinRate  float64 `json:"win_rate,omitempty"`
		PickRate float64 `json:"pick_rate,omitempty"`
		BanRate  float64 `json:"ban_rate,omitempty"`
		Kda      float64 `json:"kda,omitempty"`
		Tier     int     `json:"tier,omitempty"`
		Rank     int     `json:"rank,omitempty"`
		TierData struct {
			Tier          int `json:"tier,omitempty"`
			Rank          int `json:"rank,omitempty"`
			RankPrev      int `json:"rank_prev,omitempty"`
			RankPrevPatch int `json:"rank_prev_patch,omitempty"`
		} `json:"tier_data,omitempty"`
	} `json:"average_stats,omitempty"`
	Positions []Position `json:"positions,omitempty"`
	Roles     []Role     `json:"roles,omitempty"`
}

type ChampionInfo struct {
	Summary        ChampionSummary `json:"summary,omitempty"`
	SummonerSpells []struct {
		Stat
		Ids []int `json:"ids,omitempty"`
	} `json:"summoner_spells,omitempty"`
	CoreItems []struct {
		Stat
		Ids []int `json:"ids,omitempty"`
	} `json:"core_items,omitempty"`
	MythicItems []struct {
		Stat
		Ids []int `json:"ids,omitempty"`
	} `json:"mythic_items,omitempty"`
	Boots []struct {
		Stat
		Ids []int `json:"ids,omitempty"`
	} `json:"boots,omitempty"`
	StarterItems []struct {
		Stat
		Ids []int `json:"ids,omitempty"`
	} `json:"starter_items,omitempty"`
	LastItems []struct {
		Stat
		Ids []int `json:"ids,omitempty"`
	} `json:"last_items,omitempty"`
	RunePages []struct {
		Stat
		Id              int `json:"id,omitempty"`
		PrimaryPageId   int `json:"primary_page_id,omitempty"`
		SecondaryPageId int `json:"secondary_page_id,omitempty"`
		Builds          []struct {
			Stat
			Id               int   `json:"id,omitempty"`
			PrimaryPageId    int   `json:"primary_page_id,omitempty"`
			PrimaryRuneIds   []int `json:"primary_rune_ids,omitempty"`
			SecondaryPageId  int   `json:"secondary_page_id,omitempty"`
			SecondaryRuneIds []int `json:"secondary_rune_ids,omitempty"`
			StatModIds       []int `json:"stat_mod_ids,omitempty"`
		} `json:"builds,omitempty"`
	} `json:"rune_pages,omitempty"`
	Runes          []*Rune `json:"runes,omitempty"`
	SkillMasteries []struct {
		Stat
		Builds []Skill `json:"builds,omitempty"`
	} `json:"skill_masteries,omitempty"`
	Skills []Skill `json:"skills,omitempty"`
	Trends struct {
		TotalRank         int `json:"total_rank,omitempty"`
		TotalPositionRank int `json:"total_position_rank,omitempty"`
		Win               []struct {
			Version   string    `json:"version,omitempty"`
			Rate      float64   `json:"rate,omitempty"`
			Rank      int       `json:"rank,omitempty"`
			CreatedAt time.Time `json:"created_at,omitempty"`
		} `json:"win,omitempty"`
		Pick []struct {
			Version   string    `json:"version,omitempty"`
			Rate      float64   `json:"rate,omitempty"`
			Rank      int       `json:"rank,omitempty"`
			CreatedAt time.Time `json:"created_at,omitempty"`
		} `json:"pick,omitempty"`
		Ban []struct {
			Version   string    `json:"version,omitempty"`
			Rate      float64   `json:"rate,omitempty"`
			Rank      int       `json:"rank,omitempty"`
			CreatedAt time.Time `json:"created_at,omitempty"`
		} `json:"ban,omitempty"`
	} `json:"trends,omitempty"`
	GameLengths []struct {
		GameLength int     `json:"game_length,omitempty"`
		Rate       float64 `json:"rate,omitempty"`
		Rank       int     `json:"rank,omitempty"`
		Average    float64 `json:"average,omitempty"`
	} `json:"game_lengths,omitempty"`
	Counters []struct {
		ChampionId int `json:"champion_id,omitempty"`
		Play       int `json:"play,omitempty"`
		Win        int `json:"win,omitempty"`
	} `json:"counters,omitempty"`
}

type GetChampionInfoResponse struct {
	Data ChampionInfo `json:"data,omitempty"`
	Meta Meta         `json:"meta,omitempty"`
}

type ListChampionSummaryResponse struct {
	Data []ChampionSummary `json:"data,omitempty"`
	Meta Meta              `json:"meta,omitempty"`
}

// Stat 统计字段
type Stat struct {
	Play     int     `json:"play,omitempty"`
	Win      int     `json:"win,omitempty"`
	PickRate float64 `json:"pick_rate,omitempty"`
	WinRate  float64 `json:"-"`
	Score    float64 `json:"-"`
}

type Skill struct {
	Stat
	Order []string `json:"order,omitempty"`
}

type Rune struct {
	Stat
	Id               int   `json:"id,omitempty"`
	PrimaryPageId    int   `json:"primary_page_id,omitempty"`
	PrimaryRuneIds   []int `json:"primary_rune_ids,omitempty"`
	SecondaryPageId  int   `json:"secondary_page_id,omitempty"`
	SecondaryRuneIds []int `json:"secondary_rune_ids,omitempty"`
	StatModIds       []int `json:"stat_mod_ids,omitempty"`
}

const winRateWeight = 0.618
const pickRateWeight = 0.382

// CalcWinRate 计算胜率
func (s *Stat) CalcWinRate() {
	if s.Play > 0 {
		s.WinRate = float64(s.Win) / float64(s.Play)
	}
}

// CalcScore 计算综合得分(黄金比例权重分割, 加权平均)
func (s *Stat) CalcScore() float64 {
	s.CalcWinRate()
	s.Score = s.WinRate*winRateWeight + s.PickRate*pickRateWeight
	return s.Score
}

type IStat interface {
	CalcWinRate()
	CalcScore() float64
}

// BestEquipmentIndex 寻找最佳装备的索引
func BestEquipmentIndex[T IStat](data []T) int {
	if len(data) == 0 {
		return -1 // 返回 -1 表示数据为空
	}

	bestIndex := 0
	highestScore := 0.0

	// 计算每个装备的综合得分，并寻找最高分的索引
	for i := range data {
		score := data[i].CalcScore()
		if score > highestScore {
			highestScore = score
			bestIndex = i
		}
	}
	return bestIndex
}

// BestRune 获取最佳天赋符文
func (c *ChampionInfo) BestRune() *Rune {
	return c.Runes[BestEquipmentIndex(c.Runes)]
}
