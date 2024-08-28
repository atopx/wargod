package model

// Summoner 召唤师
type Summoner struct {
	AccountId                   int64                `json:"accountId"`
	DisplayName                 string               `json:"displayName"`
	GameName                    string               `json:"gameName"`
	InternalName                string               `json:"internalName"`
	NameChangeFlag              bool                 `json:"nameChangeFlag"`
	PercentCompleteForNextLevel int                  `json:"percentCompleteForNextLevel"`
	Privacy                     string               `json:"privacy"`
	ProfileIconId               int                  `json:"profileIconId"`
	Puuid                       string               `json:"puuid"`
	RerollPoints                SummonerRerollPoints `json:"rerollPoints"`
	SummonerId                  int64                `json:"summonerId"`
	SummonerLevel               int                  `json:"summonerLevel"`
	TagLine                     string               `json:"tagLine"`
	Unnamed                     bool                 `json:"unnamed"`
	XpSinceLastLevel            int                  `json:"xpSinceLastLevel"`
	XpUntilNextLevel            int                  `json:"xpUntilNextLevel"`
}

type SummonerRerollPoints struct {
	CurrentPoints    int `json:"currentPoints"`
	MaxRolls         int `json:"maxRolls"`
	NumberOfRolls    int `json:"numberOfRolls"`
	PointsCostToRoll int `json:"pointsCostToRoll"`
	PointsToReroll   int `json:"pointsToReroll"`
}
