package model

type Champion struct {
	// 英雄ID
	HeroId int `json:"heroId"`
	// 名称
	Name string `json:"name"`
	// 英文名
	Alias string `json:"alias"`
	// 称号
	Title string `json:"title"`
	// 角色
	Roles []string `json:"roles"`
}

type ChampAction struct {
	Id           int    `json:"id"`
	Type         string `json:"type"`
	ActorCellId  int    `json:"actorCellId"`
	ChampionId   int    `json:"championId"`
	Completed    bool   `json:"completed"`
	IsAllyAction bool   `json:"isAllyAction"`
	IsInProgress bool   `json:"isInProgress"`
}

type ChampBans struct {
	NumBans       int           `json:"numBans"`
	MyTeamBans    []interface{} `json:"myTeamBans"`
	TheirTeamBans []interface{} `json:"theirTeamBans"`
}

type MucJwt struct {
	ChannelClaim string `json:"channelClaim"`
	Domain       string `json:"domain"`
	Jwt          string `json:"jwt"`
	TargetRegion string `json:"targetRegion"`
}

type ChatDetail struct {
	MucJwtDto             MucJwt `json:"mucJwtDto"`
	MultiUserChatId       string `json:"multiUserChatId"`
	MultiUserChatPassword string `json:"multiUserChatPassword"`
}

type Timer struct {
	AdjustedTimeLeftInPhase int    `json:"adjustedTimeLeftInPhase"`
	InternalNowInEpochMs    int64  `json:"internalNowInEpochMs"`
	IsInfinite              bool   `json:"isInfinite"`
	Phase                   string `json:"phase"`
	TotalTimeInPhase        int    `json:"totalTimeInPhase"`
}

type Trade struct {
	CellId int    `json:"cellId"`
	Id     int    `json:"id"`
	State  string `json:"state"`
}

// BenchChampion 大乱斗板凳席
type BenchChampion struct {
	ChampionId int  `json:"championId"`
	IsPriority bool `json:"isPriority"`
}

type Team struct {
	AssignedPosition     string `json:"assignedPosition"`
	CellId               int    `json:"cellId"`
	ChampionId           int    `json:"championId"`
	ChampionPickIntent   int    `json:"championPickIntent"`
	NameVisibilityType   string `json:"nameVisibilityType"`
	ObfuscatedPuuid      string `json:"obfuscatedPuuid"`
	ObfuscatedSummonerId int    `json:"obfuscatedSummonerId"`
	Puuid                string `json:"puuid"`
	SelectedSkinId       int    `json:"selectedSkinId"`
	Spell1Id             int    `json:"spell1Id"`
	Spell2Id             int    `json:"spell2Id"`
	SummonerId           int64  `json:"summonerId"`
	Team                 int    `json:"team"`
	WardSkinId           int    `json:"wardSkinId"`
}

type ChampSelect struct {
	Actions              [][]ChampAction `json:"actions"`
	AllowBattleBoost     bool            `json:"allowBattleBoost"`
	AllowDuplicatePicks  bool            `json:"allowDuplicatePicks"`
	AllowLockedEvents    bool            `json:"allowLockedEvents"`
	AllowRerolling       bool            `json:"allowRerolling"`
	AllowSkinSelection   bool            `json:"allowSkinSelection"`
	Bans                 ChampBans       `json:"bans"`
	BenchChampions       []BenchChampion `json:"benchChampions"` // 大乱斗英雄板凳席
	BenchEnabled         bool            `json:"benchEnabled"`
	BoostableSkinCount   int             `json:"boostableSkinCount"`
	ChatDetails          ChatDetail      `json:"chatDetails"`
	Counter              int             `json:"counter"`
	GameId               int64           `json:"gameId"`
	HasSimultaneousBans  bool            `json:"hasSimultaneousBans"`
	HasSimultaneousPicks bool            `json:"hasSimultaneousPicks"`
	IsCustomGame         bool            `json:"isCustomGame"`
	IsSpectating         bool            `json:"isSpectating"`
	LocalPlayerCellId    int             `json:"localPlayerCellId"` // 玩家楼层
	LockedEventIndex     int             `json:"lockedEventIndex"`
	MyTeam               []Team          `json:"myTeam"` // 我方阵容
	PickOrderSwaps       []interface{}   `json:"pickOrderSwaps"`
	RecoveryCounter      int             `json:"recoveryCounter"`
	RerollsRemaining     int             `json:"rerollsRemaining"` // 剩余筛子
	SkipChampionSelect   bool            `json:"skipChampionSelect"`
	TheirTeam            []Team          `json:"theirTeam"`
	Trades               []Trade         `json:"trades"`
	Timer                Timer           `json:"timer"`
}

var ChampMap = map[int]Champion{
	1: {
		HeroId: 1,
		Title:  "黑暗之女",
		Alias:  "Annie",
		Name:   "安妮",
		Roles:  []string{"法师", "辅助"},
	},
	2: {
		HeroId: 2,
		Title:  "狂战士",
		Alias:  "Olaf",
		Name:   "奥拉夫",
		Roles:  []string{"战士", "坦克"},
	},
	3: {
		HeroId: 3,
		Title:  "正义巨像",
		Alias:  "Galio",
		Name:   "加里奥",
		Roles:  []string{"坦克", "法师"},
	},
	4: {
		HeroId: 4,
		Title:  "卡牌大师",
		Alias:  "TwistedFate",
		Name:   "崔斯特",
		Roles:  []string{"法师", "射手"},
	},
	5: {
		HeroId: 5,
		Title:  "德邦总管",
		Alias:  "XinZhao",
		Name:   "赵信",
		Roles:  []string{"战士", "坦克"},
	},
	6: {
		HeroId: 6,
		Title:  "无畏战车",
		Alias:  "Urgot",
		Name:   "厄加特",
		Roles:  []string{"战士", "坦克"},
	},
	7: {
		HeroId: 7,
		Title:  "诡术妖姬",
		Alias:  "Leblanc",
		Name:   "乐芙兰",
		Roles:  []string{"刺客", "法师"},
	},
	8: {
		HeroId: 8,
		Title:  "猩红收割者",
		Alias:  "Vladimir",
		Name:   "弗拉基米尔",
		Roles:  []string{"法师", "战士"},
	},
	9: {
		HeroId: 9,
		Title:  "远古恐惧",
		Alias:  "FiddleSticks",
		Name:   "费德提克",
		Roles:  []string{"法师", "辅助"},
	},
	10: {
		HeroId: 10,
		Title:  "正义天使",
		Alias:  "Kayle",
		Name:   "凯尔",
		Roles:  []string{"法师", "射手"},
	},
	11: {
		HeroId: 11,
		Title:  "无极剑圣",
		Alias:  "MasterYi",
		Name:   "易",
		Roles:  []string{"刺客", "战士"},
	},
	12: {
		HeroId: 12,
		Title:  "牛头酋长",
		Alias:  "Alistar",
		Name:   "阿利斯塔",
		Roles:  []string{"坦克", "辅助"},
	},
	13: {
		HeroId: 13,
		Title:  "符文法师",
		Alias:  "Ryze",
		Name:   "瑞兹",
		Roles:  []string{"法师"},
	},
	14: {
		HeroId: 14,
		Title:  "亡灵战神",
		Alias:  "Sion",
		Name:   "赛恩",
		Roles:  []string{"坦克", "战士"},
	},
	15: {
		HeroId: 15,
		Title:  "战争女神",
		Alias:  "Sivir",
		Name:   "希维尔",
		Roles:  []string{"射手"},
	},
	16: {
		HeroId: 16,
		Title:  "众星之子",
		Alias:  "Soraka",
		Name:   "索拉卡",
		Roles:  []string{"辅助", "法师"},
	},
	17: {
		HeroId: 17,
		Title:  "迅捷斥候",
		Alias:  "Teemo",
		Name:   "提莫",
		Roles:  []string{"射手", "法师"},
	},
	18: {
		HeroId: 18,
		Title:  "麦林炮手",
		Alias:  "Tristana",
		Name:   "崔丝塔娜",
		Roles:  []string{"射手", "刺客"},
	},
	19: {
		HeroId: 19,
		Title:  "祖安怒兽",
		Alias:  "Warwick",
		Name:   "沃里克",
		Roles:  []string{"战士", "坦克"},
	},
	20: {
		HeroId: 20,
		Title:  "雪原双子",
		Alias:  "Nunu",
		Name:   "努努和威朗普",
		Roles:  []string{"坦克", "法师"},
	},
	21: {
		HeroId: 21,
		Title:  "赏金猎人",
		Alias:  "MissFortune",
		Name:   "厄运小姐",
		Roles:  []string{"射手", "法师"},
	},
	22: {
		HeroId: 22,
		Title:  "寒冰射手",
		Alias:  "Ashe",
		Name:   "艾希",
		Roles:  []string{"射手", "辅助"},
	},
	23: {
		HeroId: 23,
		Title:  "蛮族之王",
		Alias:  "Tryndamere",
		Name:   "泰达米尔",
		Roles:  []string{"战士", "刺客"},
	},
	24: {
		HeroId: 24,
		Title:  "武器大师",
		Alias:  "Jax",
		Name:   "贾克斯",
		Roles:  []string{"战士"},
	},
	25: {
		HeroId: 25,
		Title:  "堕落天使",
		Alias:  "Morgana",
		Name:   "莫甘娜",
		Roles:  []string{"法师", "辅助"},
	},
	26: {
		HeroId: 26,
		Title:  "时光守护者",
		Alias:  "Zilean",
		Name:   "基兰",
		Roles:  []string{"辅助", "法师"},
	},
	27: {
		HeroId: 27,
		Title:  "炼金术士",
		Alias:  "Singed",
		Name:   "辛吉德",
		Roles:  []string{"坦克", "法师"},
	},
	28: {
		HeroId: 28,
		Title:  "痛苦之拥",
		Alias:  "Evelynn",
		Name:   "伊芙琳",
		Roles:  []string{"刺客", "法师"},
	},
	29: {
		HeroId: 29,
		Title:  "瘟疫之源",
		Alias:  "Twitch",
		Name:   "图奇",
		Roles:  []string{"射手", "刺客"},
	},
	30: {
		HeroId: 30,
		Title:  "死亡颂唱者",
		Alias:  "Karthus",
		Name:   "卡尔萨斯",
		Roles:  []string{"法师"},
	},
	31: {
		HeroId: 31,
		Title:  "虚空恐惧",
		Alias:  "Chogath",
		Name:   "科加斯",
		Roles:  []string{"坦克", "法师"},
	},
	32: {
		HeroId: 32,
		Title:  "殇之木乃伊",
		Alias:  "Amumu",
		Name:   "阿木木",
		Roles:  []string{"坦克", "辅助"},
	},
	33: {
		HeroId: 33,
		Title:  "披甲龙龟",
		Alias:  "Rammus",
		Name:   "拉莫斯",
		Roles:  []string{"坦克"},
	},
	34: {
		HeroId: 34,
		Title:  "冰晶凤凰",
		Alias:  "Anivia",
		Name:   "艾尼维亚",
		Roles:  []string{"法师"},
	},
	35: {
		HeroId: 35,
		Title:  "恶魔小丑",
		Alias:  "Shaco",
		Name:   "萨科",
		Roles:  []string{"刺客"},
	},
	36: {
		HeroId: 36,
		Title:  "祖安狂人",
		Alias:  "DrMundo",
		Name:   "蒙多医生",
		Roles:  []string{"坦克", "战士"},
	},
	37: {
		HeroId: 37,
		Title:  "琴瑟仙女",
		Alias:  "Sona",
		Name:   "娑娜",
		Roles:  []string{"辅助", "法师"},
	},
	38: {
		HeroId: 38,
		Title:  "虚空行者",
		Alias:  "Kassadin",
		Name:   "卡萨丁",
		Roles:  []string{"刺客", "法师"},
	},
	39: {
		HeroId: 39,
		Title:  "刀锋舞者",
		Alias:  "Irelia",
		Name:   "艾瑞莉娅",
		Roles:  []string{"战士", "刺客"},
	},
	40: {
		HeroId: 40,
		Title:  "风暴之怒",
		Alias:  "Janna",
		Name:   "迦娜",
		Roles:  []string{"辅助", "法师"},
	},
	41: {
		HeroId: 41,
		Title:  "海洋之灾",
		Alias:  "Gangplank",
		Name:   "普朗克",
		Roles:  []string{"战士"},
	},
	42: {
		HeroId: 42,
		Title:  "英勇投弹手",
		Alias:  "Corki",
		Name:   "库奇",
		Roles:  []string{"射手", "法师"},
	},
	43: {
		HeroId: 43,
		Title:  "天启者",
		Alias:  "Karma",
		Name:   "卡尔玛",
		Roles:  []string{"法师", "辅助"},
	},
	44: {
		HeroId: 44,
		Title:  "瓦洛兰之盾",
		Alias:  "Taric",
		Name:   "塔里克",
		Roles:  []string{"辅助", "坦克"},
	},
	45: {
		HeroId: 45,
		Title:  "邪恶小法师",
		Alias:  "Veigar",
		Name:   "维迦",
		Roles:  []string{"法师"},
	},
	48: {
		HeroId: 48,
		Title:  "巨魔之王",
		Alias:  "Trundle",
		Name:   "特朗德尔",
		Roles:  []string{"战士", "坦克"},
	},
	50: {
		HeroId: 50,
		Title:  "诺克萨斯统领",
		Alias:  "Swain",
		Name:   "斯维因",
		Roles:  []string{"法师", "辅助"},
	},
	51: {
		HeroId: 51,
		Title:  "皮城女警",
		Alias:  "Caitlyn",
		Name:   "凯特琳",
		Roles:  []string{"射手"},
	},
	53: {
		HeroId: 53,
		Title:  "蒸汽机器人",
		Alias:  "Blitzcrank",
		Name:   "布里茨",
		Roles:  []string{"坦克", "战士", "辅助"},
	},
	54: {
		HeroId: 54,
		Title:  "熔岩巨兽",
		Alias:  "Malphite",
		Name:   "墨菲特",
		Roles:  []string{"坦克", "法师"},
	},
	55: {
		HeroId: 55,
		Title:  "不祥之刃",
		Alias:  "Katarina",
		Name:   "卡特琳娜",
		Roles:  []string{"刺客", "法师"},
	},
	56: {
		HeroId: 56,
		Title:  "永恒梦魇",
		Alias:  "Nocturne",
		Name:   "魔腾",
		Roles:  []string{"战士", "刺客"},
	},
	57: {
		HeroId: 57,
		Title:  "扭曲树精",
		Alias:  "Maokai",
		Name:   "茂凯",
		Roles:  []string{"坦克", "辅助"},
	},
	58: {
		HeroId: 58,
		Title:  "荒漠屠夫",
		Alias:  "Renekton",
		Name:   "雷克顿",
		Roles:  []string{"战士", "坦克"},
	},
	59: {
		HeroId: 59,
		Title:  "德玛西亚皇子",
		Alias:  "JarvanIV",
		Name:   "嘉文四世",
		Roles:  []string{"战士", "坦克"},
	},
	60: {
		HeroId: 60,
		Title:  "蜘蛛女皇",
		Alias:  "Elise",
		Name:   "伊莉丝",
		Roles:  []string{"刺客", "法师"},
	},
	61: {
		HeroId: 61,
		Title:  "发条魔灵",
		Alias:  "Orianna",
		Name:   "奥莉安娜",
		Roles:  []string{"法师", "辅助"},
	},
	62: {
		HeroId: 62,
		Title:  "齐天大圣",
		Alias:  "MonkeyKing",
		Name:   "孙悟空",
		Roles:  []string{"战士", "坦克"},
	},
	63: {
		HeroId: 63,
		Title:  "复仇焰魂",
		Alias:  "Brand",
		Name:   "布兰德",
		Roles:  []string{"法师", "辅助"},
	},
	64: {
		HeroId: 64,
		Title:  "盲僧",
		Alias:  "LeeSin",
		Name:   "李青",
		Roles:  []string{"战士", "刺客"},
	},
	67: {
		HeroId: 67,
		Title:  "暗夜猎手",
		Alias:  "Vayne",
		Name:   "薇恩",
		Roles:  []string{"射手", "刺客"},
	},
	68: {
		HeroId: 68,
		Title:  "机械公敌",
		Alias:  "Rumble",
		Name:   "兰博",
		Roles:  []string{"战士", "法师"},
	},
	69: {
		HeroId: 69,
		Title:  "魔蛇之拥",
		Alias:  "Cassiopeia",
		Name:   "卡西奥佩娅",
		Roles:  []string{"法师"},
	},
	72: {
		HeroId: 72,
		Title:  "上古领主",
		Alias:  "Skarner",
		Name:   "斯卡纳",
		Roles:  []string{"坦克", "战士"},
	},
	74: {
		HeroId: 74,
		Title:  "大发明家",
		Alias:  "Heimerdinger",
		Name:   "黑默丁格",
		Roles:  []string{"法师", "辅助"},
	},
	75: {
		HeroId: 75,
		Title:  "沙漠死神",
		Alias:  "Nasus",
		Name:   "内瑟斯",
		Roles:  []string{"战士", "坦克"},
	},
	76: {
		HeroId: 76,
		Title:  "狂野女猎手",
		Alias:  "Nidalee",
		Name:   "奈德丽",
		Roles:  []string{"刺客", "法师"},
	},
	77: {
		HeroId: 77,
		Title:  "兽灵行者",
		Alias:  "Udyr",
		Name:   "乌迪尔",
		Roles:  []string{"战士", "坦克"},
	},
	78: {
		HeroId: 78,
		Title:  "圣锤之毅",
		Alias:  "Poppy",
		Name:   "波比",
		Roles:  []string{"坦克", "战士"},
	},
	79: {
		HeroId: 79,
		Title:  "酒桶",
		Alias:  "Gragas",
		Name:   "古拉加斯",
		Roles:  []string{"战士", "法师"},
	},
	80: {
		HeroId: 80,
		Title:  "不屈之枪",
		Alias:  "Pantheon",
		Name:   "潘森",
		Roles:  []string{"战士", "刺客"},
	},
	81: {
		HeroId: 81,
		Title:  "探险家",
		Alias:  "Ezreal",
		Name:   "伊泽瑞尔",
		Roles:  []string{"射手", "法师"},
	},
	82: {
		HeroId: 82,
		Title:  "铁铠冥魂",
		Alias:  "Mordekaiser",
		Name:   "莫德凯撒",
		Roles:  []string{"战士", "法师"},
	},
	83: {
		HeroId: 83,
		Title:  "牧魂人",
		Alias:  "Yorick",
		Name:   "约里克",
		Roles:  []string{"战士", "坦克"},
	},
	84: {
		HeroId: 84,
		Title:  "离群之刺",
		Alias:  "Akali",
		Name:   "阿卡丽",
		Roles:  []string{"刺客"},
	},
	85: {
		HeroId: 85,
		Title:  "狂暴之心",
		Alias:  "Kennen",
		Name:   "凯南",
		Roles:  []string{"法师"},
	},
	86: {
		HeroId: 86,
		Title:  "德玛西亚之力",
		Alias:  "Garen",
		Name:   "盖伦",
		Roles:  []string{"战士", "坦克"},
	},
	89: {
		HeroId: 89,
		Title:  "曙光女神",
		Alias:  "Leona",
		Name:   "蕾欧娜",
		Roles:  []string{"坦克", "辅助"},
	},
	90: {
		HeroId: 90,
		Title:  "虚空先知",
		Alias:  "Malzahar",
		Name:   "玛尔扎哈",
		Roles:  []string{"法师"},
	},
	91: {
		HeroId: 91,
		Title:  "刀锋之影",
		Alias:  "Talon",
		Name:   "泰隆",
		Roles:  []string{"刺客"},
	},
	92: {
		HeroId: 92,
		Title:  "放逐之刃",
		Alias:  "Riven",
		Name:   "锐雯",
		Roles:  []string{"战士", "刺客"},
	},
	96: {
		HeroId: 96,
		Title:  "深渊巨口",
		Alias:  "KogMaw",
		Name:   "克格莫",
		Roles:  []string{"射手", "法师"},
	},
	98: {
		HeroId: 98,
		Title:  "暮光之眼",
		Alias:  "Shen",
		Name:   "慎",
		Roles:  []string{"坦克"},
	},
	99: {
		HeroId: 99,
		Title:  "光辉女郎",
		Alias:  "Lux",
		Name:   "拉克丝",
		Roles:  []string{"法师", "辅助"},
	},
	101: {
		HeroId: 101,
		Title:  "远古巫灵",
		Alias:  "Xerath",
		Name:   "泽拉斯",
		Roles:  []string{"法师", "辅助"},
	},
	102: {
		HeroId: 102,
		Title:  "龙血武姬",
		Alias:  "Shyvana",
		Name:   "希瓦娜",
		Roles:  []string{"战士", "法师"},
	},
	103: {
		HeroId: 103,
		Title:  "九尾妖狐",
		Alias:  "Ahri",
		Name:   "阿狸",
		Roles:  []string{"法师", "刺客"},
	},
	104: {
		HeroId: 104,
		Title:  "法外狂徒",
		Alias:  "Graves",
		Name:   "格雷福斯",
		Roles:  []string{"射手"},
	},
	105: {
		HeroId: 105,
		Title:  "潮汐海灵",
		Alias:  "Fizz",
		Name:   "菲兹",
		Roles:  []string{"刺客", "战士"},
	},
	106: {
		HeroId: 106,
		Title:  "不灭狂雷",
		Alias:  "Volibear",
		Name:   "沃利贝尔",
		Roles:  []string{"战士", "坦克"},
	},
	107: {
		HeroId: 107,
		Title:  "傲之追猎者",
		Alias:  "Rengar",
		Name:   "雷恩加尔",
		Roles:  []string{"刺客", "战士"},
	},
	110: {
		HeroId: 110,
		Title:  "惩戒之箭",
		Alias:  "Varus",
		Name:   "韦鲁斯",
		Roles:  []string{"射手", "法师"},
	},
	111: {
		HeroId: 111,
		Title:  "深海泰坦",
		Alias:  "Nautilus",
		Name:   "诺提勒斯",
		Roles:  []string{"坦克", "辅助"},
	},
	112: {
		HeroId: 112,
		Title:  "机械先驱",
		Alias:  "Viktor",
		Name:   "维克托",
		Roles:  []string{"法师"},
	},
	113: {
		HeroId: 113,
		Title:  "北地之怒",
		Alias:  "Sejuani",
		Name:   "瑟庄妮",
		Roles:  []string{"坦克"},
	},
	114: {
		HeroId: 114,
		Title:  "无双剑姬",
		Alias:  "Fiora",
		Name:   "菲奥娜",
		Roles:  []string{"战士", "刺客"},
	},
	115: {
		HeroId: 115,
		Title:  "爆破鬼才",
		Alias:  "Ziggs",
		Name:   "吉格斯",
		Roles:  []string{"法师"},
	},
	117: {
		HeroId: 117,
		Title:  "仙灵女巫",
		Alias:  "Lulu",
		Name:   "璐璐",
		Roles:  []string{"辅助", "法师"},
	},
	119: {
		HeroId: 119,
		Title:  "荣耀行刑官",
		Alias:  "Draven",
		Name:   "德莱文",
		Roles:  []string{"射手"},
	},
	120: {
		HeroId: 120,
		Title:  "战争之影",
		Alias:  "Hecarim",
		Name:   "赫卡里姆",
		Roles:  []string{"战士", "坦克"},
	},
	121: {
		HeroId: 121,
		Title:  "虚空掠夺者",
		Alias:  "Khazix",
		Name:   "卡兹克",
		Roles:  []string{"刺客"},
	},
	122: {
		HeroId: 122,
		Title:  "诺克萨斯之手",
		Alias:  "Darius",
		Name:   "德莱厄斯",
		Roles:  []string{"战士", "坦克"},
	},
	126: {
		HeroId: 126,
		Title:  "未来守护者",
		Alias:  "Jayce",
		Name:   "杰斯",
		Roles:  []string{"射手", "战士"},
	},
	127: {
		HeroId: 127,
		Title:  "冰霜女巫",
		Alias:  "Lissandra",
		Name:   "丽桑卓",
		Roles:  []string{"法师"},
	},
	131: {
		HeroId: 131,
		Title:  "皎月女神",
		Alias:  "Diana",
		Name:   "黛安娜",
		Roles:  []string{"战士", "刺客"},
	},
	133: {
		HeroId: 133,
		Title:  "德玛西亚之翼",
		Alias:  "Quinn",
		Name:   "奎因",
		Roles:  []string{"射手", "刺客"},
	},
	134: {
		HeroId: 134,
		Title:  "暗黑元首",
		Alias:  "Syndra",
		Name:   "辛德拉",
		Roles:  []string{"法师"},
	},
	136: {
		HeroId: 136,
		Title:  "铸星龙王",
		Alias:  "AurelionSol",
		Name:   "奥瑞利安索尔",
		Roles:  []string{"法师"},
	},
	141: {
		HeroId: 141,
		Title:  "影流之镰",
		Alias:  "Kayn",
		Name:   "凯隐",
		Roles:  []string{"战士", "刺客"},
	},
	142: {
		HeroId: 142,
		Title:  "暮光星灵",
		Alias:  "Zoe",
		Name:   "佐伊",
		Roles:  []string{"法师"},
	},
	143: {
		HeroId: 143,
		Title:  "荆棘之兴",
		Alias:  "Zyra",
		Name:   "婕拉",
		Roles:  []string{"法师", "辅助"},
	},
	145: {
		HeroId: 145,
		Title:  "虚空之女",
		Alias:  "Kaisa",
		Name:   "卡莎",
		Roles:  []string{"射手", "法师"},
	},
	147: {
		HeroId: 147,
		Title:  "星籁歌姬",
		Alias:  "Seraphine",
		Name:   "萨勒芬妮",
		Roles:  []string{"辅助", "法师"},
	},
	150: {
		HeroId: 150,
		Title:  "迷失之牙",
		Alias:  "Gnar",
		Name:   "纳尔",
		Roles:  []string{"战士", "坦克"},
	},
	154: {
		HeroId: 154,
		Title:  "生化魔人",
		Alias:  "Zac",
		Name:   "扎克",
		Roles:  []string{"坦克", "战士"},
	},
	157: {
		HeroId: 157,
		Title:  "疾风剑豪",
		Alias:  "Yasuo",
		Name:   "亚索",
		Roles:  []string{"战士", "刺客"},
	},
	161: {
		HeroId: 161,
		Title:  "虚空之眼",
		Alias:  "Velkoz",
		Name:   "维克兹",
		Roles:  []string{"法师", "辅助"},
	},
	163: {
		HeroId: 163,
		Title:  "岩雀",
		Alias:  "Taliyah",
		Name:   "塔莉垭",
		Roles:  []string{"法师", "辅助"},
	},
	164: {
		HeroId: 164,
		Title:  "青钢影",
		Alias:  "Camille",
		Name:   "卡蜜尔",
		Roles:  []string{"战士", "刺客"},
	},
	166: {
		HeroId: 166,
		Title:  "影哨",
		Alias:  "Akshan",
		Name:   "阿克尚",
		Roles:  []string{"射手", "刺客"},
	},
	200: {
		HeroId: 200,
		Title:  "虚空女皇",
		Alias:  "Belveth",
		Name:   "卑尔维斯",
		Roles:  []string{"战士"},
	},
	201: {
		HeroId: 201,
		Title:  "弗雷尔卓德之心",
		Alias:  "Braum",
		Name:   "布隆",
		Roles:  []string{"坦克", "辅助"},
	},
	202: {
		HeroId: 202,
		Title:  "戏命师",
		Alias:  "Jhin",
		Name:   "烬",
		Roles:  []string{"射手", "法师"},
	},
	203: {
		HeroId: 203,
		Title:  "永猎双子",
		Alias:  "Kindred",
		Name:   "千珏",
		Roles:  []string{"射手"},
	},
	221: {
		HeroId: 221,
		Title:  "祖安花火",
		Alias:  "Zeri",
		Name:   "泽丽",
		Roles:  []string{"射手"},
	},
	222: {
		HeroId: 222,
		Title:  "暴走萝莉",
		Alias:  "Jinx",
		Name:   "金克丝",
		Roles:  []string{"射手"},
	},
	223: {
		HeroId: 223,
		Title:  "河流之王",
		Alias:  "TahmKench",
		Name:   "塔姆",
		Roles:  []string{"坦克", "辅助"},
	},
	233: {
		HeroId: 233,
		Title:  "狂厄蔷薇",
		Alias:  "Briar",
		Name:   "贝蕾亚",
		Roles:  []string{"战士", "刺客"},
	},
	234: {
		HeroId: 234,
		Title:  "破败之王",
		Alias:  "Viego",
		Name:   "佛耶戈",
		Roles:  []string{"战士", "刺客"},
	},
	235: {
		HeroId: 235,
		Title:  "涤魂圣枪",
		Alias:  "Senna",
		Name:   "赛娜",
		Roles:  []string{"辅助", "射手"},
	},
	236: {
		HeroId: 236,
		Title:  "圣枪游侠",
		Alias:  "Lucian",
		Name:   "卢锡安",
		Roles:  []string{"射手", "刺客"},
	},
	238: {
		HeroId: 238,
		Title:  "影流之主",
		Alias:  "Zed",
		Name:   "劫",
		Roles:  []string{"刺客"},
	},
	240: {
		HeroId: 240,
		Title:  "暴怒骑士",
		Alias:  "Kled",
		Name:   "克烈",
		Roles:  []string{"战士"},
	},
	245: {
		HeroId: 245,
		Title:  "时间刺客",
		Alias:  "Ekko",
		Name:   "艾克",
		Roles:  []string{"刺客", "法师"},
	},
	246: {
		HeroId: 246,
		Title:  "元素女皇",
		Alias:  "Qiyana",
		Name:   "奇亚娜",
		Roles:  []string{"刺客"},
	},
	254: {
		HeroId: 254,
		Title:  "皮城执法官",
		Alias:  "Vi",
		Name:   "蔚",
		Roles:  []string{"战士", "刺客"},
	},
	266: {
		HeroId: 266,
		Title:  "暗裔剑魔",
		Alias:  "Aatrox",
		Name:   "亚托克斯",
		Roles:  []string{"战士"},
	},
	267: {
		HeroId: 267,
		Title:  "唤潮鲛姬",
		Alias:  "Nami",
		Name:   "娜美",
		Roles:  []string{"辅助", "法师"},
	},
	268: {
		HeroId: 268,
		Title:  "沙漠皇帝",
		Alias:  "Azir",
		Name:   "阿兹尔",
		Roles:  []string{"法师", "射手"},
	},
	350: {
		HeroId: 350,
		Title:  "魔法猫咪",
		Alias:  "Yuumi",
		Name:   "悠米",
		Roles:  []string{"辅助", "法师"},
	},
	360: {
		HeroId: 360,
		Title:  "沙漠玫瑰",
		Alias:  "Samira",
		Name:   "莎弥拉",
		Roles:  []string{"射手", "刺客"},
	},
	412: {
		HeroId: 412,
		Title:  "魂锁典狱长",
		Alias:  "Thresh",
		Name:   "锤石",
		Roles:  []string{"辅助", "坦克"},
	},
	420: {
		HeroId: 420,
		Title:  "海兽祭司",
		Alias:  "Illaoi",
		Name:   "俄洛伊",
		Roles:  []string{"战士", "坦克"},
	},
	421: {
		HeroId: 421,
		Title:  "虚空遁地兽",
		Alias:  "RekSai",
		Name:   "雷克塞",
		Roles:  []string{"战士", "坦克"},
	},
	427: {
		HeroId: 427,
		Title:  "翠神",
		Alias:  "Ivern",
		Name:   "艾翁",
		Roles:  []string{"辅助", "法师"},
	},
	429: {
		HeroId: 429,
		Title:  "复仇之矛",
		Alias:  "Kalista",
		Name:   "卡莉丝塔",
		Roles:  []string{"射手"},
	},
	432: {
		HeroId: 432,
		Title:  "星界游神",
		Alias:  "Bard",
		Name:   "巴德",
		Roles:  []string{"辅助", "法师"},
	},
	497: {
		HeroId: 497,
		Title:  "幻翎",
		Alias:  "Rakan",
		Name:   "洛",
		Roles:  []string{"辅助"},
	},
	498: {
		HeroId: 498,
		Title:  "逆羽",
		Alias:  "Xayah",
		Name:   "霞",
		Roles:  []string{"射手"},
	},
	516: {
		HeroId: 516,
		Title:  "山隐之焰",
		Alias:  "Ornn",
		Name:   "奥恩",
		Roles:  []string{"坦克"},
	},
	517: {
		HeroId: 517,
		Title:  "解脱者",
		Alias:  "Sylas",
		Name:   "塞拉斯",
		Roles:  []string{"法师", "刺客"},
	},
	518: {
		HeroId: 518,
		Title:  "万花通灵",
		Alias:  "Neeko",
		Name:   "妮蔻",
		Roles:  []string{"法师", "辅助"},
	},
	523: {
		HeroId: 523,
		Title:  "残月之肃",
		Alias:  "Aphelios",
		Name:   "厄斐琉斯",
		Roles:  []string{"射手"},
	},
	526: {
		HeroId: 526,
		Title:  "镕铁少女",
		Alias:  "Rell",
		Name:   "芮尔",
		Roles:  []string{"坦克", "辅助"},
	},
	555: {
		HeroId: 555,
		Title:  "血港鬼影",
		Alias:  "Pyke",
		Name:   "派克",
		Roles:  []string{"辅助", "刺客"},
	},
	711: {
		HeroId: 711,
		Title:  "愁云使者",
		Alias:  "Vex",
		Name:   "薇古丝",
		Roles:  []string{"法师"},
	},
	777: {
		HeroId: 777,
		Title:  "封魔剑魂",
		Alias:  "Yone",
		Name:   "永恩",
		Roles:  []string{"战士", "刺客"},
	},
	875: {
		HeroId: 875,
		Title:  "腕豪",
		Alias:  "Sett",
		Name:   "瑟提",
		Roles:  []string{"战士", "坦克"},
	},
	876: {
		HeroId: 876,
		Title:  "含羞蓓蕾",
		Alias:  "Lillia",
		Name:   "莉莉娅",
		Roles:  []string{"战士", "法师"},
	},
	887: {
		HeroId: 887,
		Title:  "灵罗娃娃",
		Alias:  "Gwen",
		Name:   "格温",
		Roles:  []string{"战士"},
	},
	888: {
		HeroId: 888,
		Title:  "炼金男爵",
		Alias:  "Renata",
		Name:   "烈娜塔 · 戈拉斯克",
		Roles:  []string{"辅助", "法师"},
	},
	893: {
		HeroId: 893,
		Title:  "双界灵兔",
		Alias:  "Aurora",
		Name:   "阿萝拉",
		Roles:  []string{"法师", "刺客"},
	},
	895: {
		HeroId: 895,
		Title:  "不羁之悦",
		Alias:  "Nilah",
		Name:   "尼菈",
		Roles:  []string{"战士", "刺客"},
	},
	897: {
		HeroId: 897,
		Title:  "纳祖芒荣耀",
		Alias:  "KSante",
		Name:   "奎桑提",
		Roles:  []string{"坦克", "战士"},
	},
	901: {
		HeroId: 901,
		Title:  "炽炎雏龙",
		Alias:  "Smolder",
		Name:   "斯莫德",
		Roles:  []string{"射手", "法师"},
	},
	902: {
		HeroId: 902,
		Title:  "明烛",
		Alias:  "Milio",
		Name:   "米利欧",
		Roles:  []string{"辅助", "法师"},
	},
	910: {
		HeroId: 910,
		Title:  "异画师",
		Alias:  "Hwei",
		Name:   "彗",
		Roles:  []string{"法师", "辅助"},
	},
	950: {
		HeroId: 950,
		Title:  "百裂冥犬",
		Alias:  "Naafiri",
		Name:   "纳亚菲利",
		Roles:  []string{"刺客", "战士"},
	},
}
