package model

import "time"

// GameSummary 战绩列表
type GameSummary struct {
	AccountId  int64      `json:"accountId"` // 当前战绩列表所属玩家的账户ID
	PlatformId string     `json:"platformId"`
	Games      GameRecord `json:"games"`
}

type GameRecord struct {
	GameBeginDate  string     `json:"gameBeginDate"`
	GameCount      int        `json:"gameCount"`      // 可查询的对局总数
	GameEndDate    string     `json:"gameEndDate"`    // 游戏结束时间
	GameIndexBegin int        `json:"gameIndexBegin"` // 返回的对局中第一场对局的序号
	GameIndexEnd   int        `json:"gameIndexEnd"`   // 返回对局中最后一场对局的序号
	Games          []GameInfo `json:"games"`
}

// GameParticipantIdentity 玩家身份信息列表
type GameParticipantIdentity struct {
	ParticipantId int `json:"participantId"` // 对局玩家序号 0-9
	Player        struct {
		AccountId         int    `json:"accountId"`
		CurrentAccountId  int    `json:"currentAccountId"`
		CurrentPlatformId string `json:"currentPlatformId"`
		GameName          string `json:"gameName"`
		MatchHistoryUri   string `json:"matchHistoryUri"`
		PlatformId        string `json:"platformId"`
		ProfileIcon       int    `json:"profileIcon"`
		Puuid             string `json:"puuid"`        // 召唤师PUUID
		SummonerId        int64  `json:"summonerId"`   // 召唤师ID
		SummonerName      string `json:"summonerName"` // 召唤师名称
		TagLine           string `json:"tagLine"`
	} `json:"player"`
}

// GameStats 统计信息
type GameStats struct {
	Assists                         int  `json:"assists"`                   // 助攻数
	CausedEarlySurrender            bool `json:"causedEarlySurrender"`      // 是否提前投降
	ChampLevel                      int  `json:"champLevel"`                // 结束时的英雄等级
	CombatPlayerScore               int  `json:"combatPlayerScore"`         //
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`   // 对战略点的伤害
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`      // 对防御塔的伤害
	DamageSelfMitigated             int  `json:"damageSelfMitigated"`       //
	Deaths                          int  `json:"deaths"`                    // 死亡数
	DoubleKills                     int  `json:"doubleKills"`               // 双杀次数
	EarlySurrenderAccomplice        bool `json:"earlySurrenderAccomplice"`  // 参与提前投降
	FirstBloodAssist                bool `json:"firstBloodAssist"`          // 一血辅助
	FirstBloodKill                  bool `json:"firstBloodKill"`            // 获得一血
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`      // 助攻一血水晶
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`        // 一血水晶
	FirstTowerAssist                bool `json:"firstTowerAssist"`          // 辅助一血塔
	FirstTowerKill                  bool `json:"firstTowerKill"`            // 一血塔
	GameEndedInEarlySurrender       bool `json:"gameEndedInEarlySurrender"` // 游戏以提前投降结束
	GameEndedInSurrender            bool `json:"gameEndedInSurrender"`      // 游戏以投降结束
	GoldEarned                      int  `json:"goldEarned"`                // 获取的金币
	GoldSpent                       int  `json:"goldSpent"`                 // 花费的金币
	InhibitorKills                  int  `json:"inhibitorKills"`            // 摧毁水晶
	Item0                           int  `json:"item0"`                     // 装备1的ID
	Item1                           int  `json:"item1"`                     // 装备2的ID
	Item2                           int  `json:"item2"`                     // 装备3的ID
	Item3                           int  `json:"item3"`                     // 装备4的ID
	Item4                           int  `json:"item4"`                     // 装备5的ID
	Item5                           int  `json:"item5"`                     // 装备6的ID
	Item6                           int  `json:"item6"`                     // 装备7的ID
	KillingSprees                   int  `json:"killingSprees"`
	Kills                           int  `json:"kills"`                 // 击杀数
	LargestCriticalStrike           int  `json:"largestCriticalStrike"` // 最大暴击伤害
	LargestKillingSpree             int  `json:"largestKillingSpree"`   // 最大连杀
	LargestMultiKill                int  `json:"largestMultiKill"`      // 最大多杀
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"` // 对英雄的魔法伤害
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`          // 承受的魔法伤害
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`        // 击杀野怪数
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	ParticipantId                   int  `json:"participantId"`
	PentaKills                      int  `json:"pentaKills"` // 五杀次数
	Perk0                           int  `json:"perk0"`      // perk0~perk5: 携带的5个符文ID
	Perk0Var1                       int  `json:"perk0Var1"`  // perk[x]Var[y]: 第x符文的y项数据
	Perk0Var2                       int  `json:"perk0Var2"`
	Perk0Var3                       int  `json:"perk0Var3"`
	Perk1                           int  `json:"perk1"`
	Perk1Var1                       int  `json:"perk1Var1"`
	Perk1Var2                       int  `json:"perk1Var2"`
	Perk1Var3                       int  `json:"perk1Var3"`
	Perk2                           int  `json:"perk2"`
	Perk2Var1                       int  `json:"perk2Var1"`
	Perk2Var2                       int  `json:"perk2Var2"`
	Perk2Var3                       int  `json:"perk2Var3"`
	Perk3                           int  `json:"perk3"`
	Perk3Var1                       int  `json:"perk3Var1"`
	Perk3Var2                       int  `json:"perk3Var2"`
	Perk3Var3                       int  `json:"perk3Var3"`
	Perk4                           int  `json:"perk4"`
	Perk4Var1                       int  `json:"perk4Var1"`
	Perk4Var2                       int  `json:"perk4Var2"`
	Perk4Var3                       int  `json:"perk4Var3"`
	Perk5                           int  `json:"perk5"`
	Perk5Var1                       int  `json:"perk5Var1"`
	Perk5Var2                       int  `json:"perk5Var2"`
	Perk5Var3                       int  `json:"perk5Var3"`
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`               // 主系符文
	PerkSubStyle                    int  `json:"perkSubStyle"`                   // 副系符文
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`            // 物理伤害总和
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"` // 对英雄的物理伤害
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`            // 承受的物理伤害
	PlayerAugment1                  int  `json:"playerAugment1"`
	PlayerAugment2                  int  `json:"playerAugment2"`
	PlayerAugment3                  int  `json:"playerAugment3"`
	PlayerAugment4                  int  `json:"playerAugment4"`
	PlayerAugment5                  int  `json:"playerAugment5"`
	PlayerAugment6                  int  `json:"playerAugment6"`
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore6                    int  `json:"playerScore6"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore8                    int  `json:"playerScore8"`
	PlayerScore9                    int  `json:"playerScore9"`
	PlayerSubteamId                 int  `json:"playerSubteamId"`
	QuadraKills                     int  `json:"quadraKills"`            // 四杀数
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"` // 买眼数
	SubteamPlacement                int  `json:"subteamPlacement"`
	TeamEarlySurrendered            bool `json:"teamEarlySurrendered"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`            // 造成的总伤害
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"` // 对英雄的总伤害
	TotalDamageTaken                int  `json:"totalDamageTaken"`            // 承受的总伤害
	TotalHeal                       int  `json:"totalHeal"`                   // 总治疗
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`          // 补兵数量
	TotalPlayerScore                int  `json:"totalPlayerScore"`            // 玩家总得分
	TotalScoreRank                  int  `json:"totalScoreRank"`              // 总得分排名
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TripleKills                     int  `json:"tripleKills"`                // 三杀次数
	TrueDamageDealt                 int  `json:"trueDamageDealt"`            // 造成的真实伤害
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"` // 对英雄造成的真实伤害
	TrueDamageTaken                 int  `json:"trueDamageTaken"`            // 承受的真实伤害
	TurretKills                     int  `json:"turretKills"`
	UnrealKills                     int  `json:"unrealKills"`
	VisionScore                     int  `json:"visionScore"`             // 视野得分
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"` // 购买控制守卫
	WardsKilled                     int  `json:"wardsKilled"`             // 守卫被杀数
	WardsPlaced                     int  `json:"wardsPlaced"`             // 放置守卫数
	Win                             bool `json:"win"`                     // 胜负
}

type GameInfo struct {
	EndOfGameResult       string                    `json:"endOfGameResult"`  // 对局结束状态
	GameCreation          int64                     `json:"gameCreation"`     // 对局创建时的时间戳
	GameCreationDate      time.Time                 `json:"gameCreationDate"` // 对局创建时间
	GameDuration          int                       `json:"gameDuration"`     // 对局时长，单位秒
	GameId                int64                     `json:"gameId"`           // 对局ID，用于后续查询单场详情
	GameMode              string                    `json:"gameMode"`         // 游戏模式: ARAM/CLASSIC
	GameType              string                    `json:"gameType"`         // 游戏类型: MATCHED_GAME / RANKED_GAME
	GameVersion           string                    `json:"gameVersion"`      // 对局游戏版本号: 14.16.611.2424
	MapId                 int                       `json:"mapId"`            // 地图ID
	ParticipantIdentities []GameParticipantIdentity `json:"participantIdentities"`
	Participants          []GameParticipant         `json:"participants"`
	PlatformId            string                    `json:"platformId"`
	QueueId               int                       `json:"queueId"`
	SeasonId              int                       `json:"seasonId"`
	Teams                 []GameTeam                `json:"teams"`
}

type GameTimelineItem struct {
	AdditionalProp1 int `json:"additionalProp1"`
	AdditionalProp2 int `json:"additionalProp2"`
	AdditionalProp3 int `json:"additionalProp3"`
}

// GameTimeline 游戏时间线
type GameTimeline struct {
	Lane                        string           `json:"lane"`
	Role                        string           `json:"role"`
	ParticipantId               int              `json:"participantId"`
	CreepsPerMinDeltas          GameTimelineItem `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas          GameTimelineItem `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas GameTimelineItem `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     GameTimelineItem `json:"damageTakenPerMinDeltas"`
	GoldPerMinDeltas            GameTimelineItem `json:"goldPerMinDeltas"`
	XpDiffPerMinDeltas          GameTimelineItem `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas              GameTimelineItem `json:"xpPerMinDeltas"`
}

// GameParticipant 玩家游戏详细信息列表
type GameParticipant struct {
	TeamId                    int          `json:"teamId"`                    // 红蓝方阵营 (红:100)或(蓝:200)
	ChampionId                int          `json:"championId"`                // 英雄ID
	HighestAchievedSeasonTier string       `json:"highestAchievedSeasonTier"` // 玩家达到过的最高段位
	ParticipantId             int          `json:"participantId"`             // 对局玩家序号，从0~9
	Spell1Id                  int          `json:"spell1Id"`                  // 召唤师技能1ID
	Spell2Id                  int          `json:"spell2Id"`                  // 召唤师技能2ID
	Stats                     GameStats    `json:"stats"`
	Timeline                  GameTimeline `json:"timeline"`
}

type GameBan struct {
	ChampionId int `json:"championId"` // ban的英雄ID
	PickTurn   int `json:"pickTurn"`   // ban的次序
}

type GameTeam struct {
	TeamId               int       `json:"teamId"`     // 红蓝方阵营 (红:100)或(蓝:200)
	Bans                 []GameBan `json:"bans"`       // BAN的英雄
	BaronKills           int       `json:"baronKills"` // 大龙击杀数
	DominionVictoryScore int       `json:"dominionVictoryScore"`
	DragonKills          int       `json:"dragonKills"`    // 小龙击杀数
	FirstBaron           bool      `json:"firstBaron"`     // 首杀大龙
	FirstBlood           bool      `json:"firstBlood"`     // 一血
	FirstDargon          bool      `json:"firstDargon"`    //
	FirstInhibitor       bool      `json:"firstInhibitor"` // 摧毁首个水晶
	FirstTower           bool      `json:"firstTower"`     // 一血塔
	HordeKills           int       `json:"hordeKills"`
	InhibitorKills       int       `json:"inhibitorKills"`  // 摧毁水晶数
	RiftHeraldKills      int       `json:"riftHeraldKills"` //
	TowerKills           int       `json:"towerKills"`      // 推塔数
	VilemawKills         int       `json:"vilemawKills"`
	Win                  string    `json:"win"` //  Fail | Win
}
