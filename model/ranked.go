package model

import "wargod/types"

type Ranked struct {
	PreviousSeasonSplitPoints             int                               `json:"previousSeasonSplitPoints"`
	EarnedRegaliaRewardIds                []string                          `json:"earnedRegaliaRewardIds"`
	HighestPreviousSeasonAchievedTier     types.Tier                        `json:"highestPreviousSeasonAchievedTier"`     // 上赛季最高段位
	HighestPreviousSeasonAchievedDivision types.TierDivision                `json:"highestPreviousSeasonAchievedDivision"` // 上赛季最高段位等级
	HighestPreviousSeasonEndDivision      string                            `json:"highestPreviousSeasonEndDivision"`      // 上赛季结算段位等级
	HighestPreviousSeasonEndTier          string                            `json:"highestPreviousSeasonEndTier"`          // 算个赛季结算段位
	HighestRankedEntry                    RankedEntry                       `json:"highestRankedEntry"`                    // 单双最高段位
	HighestRankedEntrySR                  RankedEntry                       `json:"highestRankedEntrySR"`                  // 组排最高段位
	QueueMap                              map[types.RankedQueue]RankedEntry `json:"queueMap"`                              // 排位各队列数据
	Queues                                []RankedEntry                     `json:"queues"`                                // 所有队列数据
	RankedRegaliaLevel                    int                               `json:"rankedRegaliaLevel"`
	Seasons                               map[types.RankedQueue]RankedEntry `json:"seasons"`
	SplitsProgress                        struct {
		AdditionalProp1 int `json:"additionalProp1"`
		AdditionalProp2 int `json:"additionalProp2"`
		AdditionalProp3 int `json:"additionalProp3"`
	} `json:"splitsProgress"`
}

type RankedSeason struct {
	CurrentSeasonEnd int `json:"currentSeasonEnd"`
	CurrentSeasonId  int `json:"currentSeasonId"`
	NextSeasonStart  int `json:"nextSeasonStart"`
}

type RankedEntry struct {
	Losses                         int                `json:"losses"`                         // 败场数量
	Wins                           int                `json:"wins"`                           // 胜场数
	LeaguePoints                   int                `json:"leaguePoints"`                   // 胜点
	Division                       types.TierDivision `json:"division"`                       // 段位等级
	HighestDivision                types.TierDivision `json:"highestDivision"`                // 最高段位等级
	Tier                           types.Tier         `json:"tier"`                           // 当前段位
	HighestTier                    types.Tier         `json:"highestTier"`                    // 最高段位
	PreviousSeasonAchievedDivision types.TierDivision `json:"previousSeasonAchievedDivision"` // 上赛季段位等级
	PreviousSeasonHighestDivision  types.TierDivision `json:"previousSeasonHighestDivision"`  // 上赛季最高段位等级
	PreviousSeasonEndDivision      types.TierDivision `json:"previousSeasonEndDivision"`      // 上赛季结算段位等级
	PreviousSeasonAchievedTier     types.Tier         `json:"previousSeasonAchievedTier"`     // 上赛季段位
	PreviousSeasonHighestTier      types.Tier         `json:"previousSeasonHighestTier"`      // 上赛季最高段位
	PreviousSeasonEndTier          types.Tier         `json:"previousSeasonEndTier"`          // 上赛季结算段位
	QueueType                      string             `json:"queueType"`                      // 队列类型
	RatedRating                    int                `json:"ratedRating"`                    // 额定评级
	RatedTier                      string             `json:"ratedTier"`                      // 评级段位
	IsProvisional                  bool               `json:"isProvisional"`                  // 临时的
	MiniSeriesProgress             string             `json:"miniSeriesProgress"`
	ProvisionalGameThreshold       int                `json:"provisionalGameThreshold"`
	ProvisionalGamesRemaining      int                `json:"provisionalGamesRemaining"`
	Warnings                       struct {
		DaysUntilDecay                   int  `json:"daysUntilDecay"`
		DemotionWarning                  int  `json:"demotionWarning"`
		DisplayDecayWarning              bool `json:"displayDecayWarning"`
		TimeUntilInactivityStatusChanges int  `json:"timeUntilInactivityStatusChanges"`
	} `json:"warnings"`
}
