package types

type RankedQueue string

const (
	RankedSolo5x5 RankedQueue = "RANKED_SOLO_5x5" // 单双
	RankedFlexSr  RankedQueue = "RANKED_FLEX_SR"  // 灵活组排
	RankedTft     RankedQueue = "RANKED_TFT"      // 云顶之弈
)

type Tier string

const (
	TierNone        Tier = "NONE"        // 无段位
	TierIron        Tier = "IRON"        // 黑铁
	TierBronze      Tier = "BRONZE"      // 青铜
	TierSilver      Tier = "SILVER"      // 白银
	TierGold        Tier = "GOLD"        // 黄金
	TierPlatinum    Tier = "PLATINUM"    // 铂金
	TierEmerald     Tier = "EMERALD"     // 翡翠
	TierDiamond     Tier = "DIAMOND"     // 钻石
	TierMaster      Tier = "MASTER"      // 大师
	TierGrandmaster Tier = "GRANDMASTER" // 宗师
	TierChallenger  Tier = "CHALLENGER"  // 王者
)

type TierDivision string

const (
	TierDivision0 TierDivision = "NA"
	TierDivision1 TierDivision = "I"
	TierDivision2 TierDivision = "II"
	TierDivision3 TierDivision = "III"
	TierDivision4 TierDivision = "IV"
	TierDivision5 TierDivision = "V"
)

var (
	AllQueues = []RankedQueue{
		RankedSolo5x5,
		RankedFlexSr,
		RankedTft,
	}

	AllTiers = []Tier{
		TierNone,
		TierIron,
		TierBronze,
		TierSilver,
		TierGold,
		TierPlatinum,
		TierEmerald,
		TierDiamond,
		TierMaster,
		TierGrandmaster,
		TierChallenger,
	}

	AllTierDivision = []TierDivision{
		TierDivision0,
		TierDivision1,
		TierDivision2,
		TierDivision3,
		TierDivision4,
		TierDivision5,
	}
)
