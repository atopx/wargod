package model

type Lobby struct {
	CanStartActivity bool            `json:"canStartActivity"`
	GameConfig       LobbyGameConfig `json:"gameConfig"`
	//Invitations      []struct {
	//	InvitationId   string `json:"invitationId"`
	//	InvitationType string `json:"invitationType"`
	//	State          string `json:"state"`
	//	Timestamp      string `json:"timestamp"`
	//	ToSummonerId   int64  `json:"toSummonerId"`
	//	ToSummonerName string `json:"toSummonerName"`
	//} `json:"invitations"`
	//LocalMember struct {
	//	AllowedChangeActivity         bool          `json:"allowedChangeActivity"`
	//	AllowedInviteOthers           bool          `json:"allowedInviteOthers"`
	//	AllowedKickOthers             bool          `json:"allowedKickOthers"`
	//	AllowedStartActivity          bool          `json:"allowedStartActivity"`
	//	AllowedToggleInvite           bool          `json:"allowedToggleInvite"`
	//	AutoFillEligible              bool          `json:"autoFillEligible"`
	//	AutoFillProtectedForPromos    bool          `json:"autoFillProtectedForPromos"`
	//	AutoFillProtectedForRemedy    bool          `json:"autoFillProtectedForRemedy"`
	//	AutoFillProtectedForSoloing   bool          `json:"autoFillProtectedForSoloing"`
	//	AutoFillProtectedForStreaking bool          `json:"autoFillProtectedForStreaking"`
	//	BotChampionId                 int           `json:"botChampionId"`
	//	BotDifficulty                 string        `json:"botDifficulty"`
	//	BotId                         string        `json:"botId"`
	//	BotPosition                   string        `json:"botPosition"`
	//	BotUuid                       string        `json:"botUuid"`
	//	FirstPositionPreference       string        `json:"firstPositionPreference"`
	//	IntraSubteamPosition          interface{}   `json:"intraSubteamPosition"`
	//	IsBot                         bool          `json:"isBot"`
	//	IsLeader                      bool          `json:"isLeader"`
	//	IsSpectator                   bool          `json:"isSpectator"`
	//	PlayerSlots                   []interface{} `json:"playerSlots"`
	//	Puuid                         string        `json:"puuid"`
	//	QuickplayPlayerState          interface{}   `json:"quickplayPlayerState"`
	//	Ready                         bool          `json:"ready"`
	//	SecondPositionPreference      string        `json:"secondPositionPreference"`
	//	ShowGhostedBanner             bool          `json:"showGhostedBanner"`
	//	StrawberryMapId               interface{}   `json:"strawberryMapId"`
	//	SubteamIndex                  interface{}   `json:"subteamIndex"`
	//	SummonerIconId                int           `json:"summonerIconId"`
	//	SummonerId                    int64         `json:"summonerId"`
	//	SummonerInternalName          string        `json:"summonerInternalName"`
	//	SummonerLevel                 int           `json:"summonerLevel"`
	//	SummonerName                  string        `json:"summonerName"`
	//	TeamId                        int           `json:"teamId"`
	//	TftNPEQueueBypass             bool          `json:"tftNPEQueueBypass"`
	//} `json:"localMember"`
	//Members []struct {
	//	AllowedChangeActivity         bool          `json:"allowedChangeActivity"`
	//	AllowedInviteOthers           bool          `json:"allowedInviteOthers"`
	//	AllowedKickOthers             bool          `json:"allowedKickOthers"`
	//	AllowedStartActivity          bool          `json:"allowedStartActivity"`
	//	AllowedToggleInvite           bool          `json:"allowedToggleInvite"`
	//	AutoFillEligible              bool          `json:"autoFillEligible"`
	//	AutoFillProtectedForPromos    bool          `json:"autoFillProtectedForPromos"`
	//	AutoFillProtectedForRemedy    bool          `json:"autoFillProtectedForRemedy"`
	//	AutoFillProtectedForSoloing   bool          `json:"autoFillProtectedForSoloing"`
	//	AutoFillProtectedForStreaking bool          `json:"autoFillProtectedForStreaking"`
	//	BotChampionId                 int           `json:"botChampionId"`
	//	BotDifficulty                 string        `json:"botDifficulty"`
	//	BotId                         string        `json:"botId"`
	//	BotPosition                   string        `json:"botPosition"`
	//	BotUuid                       string        `json:"botUuid"`
	//	FirstPositionPreference       string        `json:"firstPositionPreference"`
	//	IntraSubteamPosition          interface{}   `json:"intraSubteamPosition"`
	//	IsBot                         bool          `json:"isBot"`
	//	IsLeader                      bool          `json:"isLeader"`
	//	IsSpectator                   bool          `json:"isSpectator"`
	//	PlayerSlots                   []interface{} `json:"playerSlots"`
	//	Puuid                         string        `json:"puuid"`
	//	QuickplayPlayerState          interface{}   `json:"quickplayPlayerState"`
	//	Ready                         bool          `json:"ready"`
	//	SecondPositionPreference      string        `json:"secondPositionPreference"`
	//	ShowGhostedBanner             bool          `json:"showGhostedBanner"`
	//	StrawberryMapId               interface{}   `json:"strawberryMapId"`
	//	SubteamIndex                  interface{}   `json:"subteamIndex"`
	//	SummonerIconId                int           `json:"summonerIconId"`
	//	SummonerId                    int64         `json:"summonerId"`
	//	SummonerInternalName          string        `json:"summonerInternalName"`
	//	SummonerLevel                 int           `json:"summonerLevel"`
	//	SummonerName                  string        `json:"summonerName"`
	//	TeamId                        int           `json:"teamId"`
	//	TftNPEQueueBypass             bool          `json:"tftNPEQueueBypass"`
	//} `json:"members"`
	//MucJwtDto struct {
	//	ChannelClaim string `json:"channelClaim"`
	//	Domain       string `json:"domain"`
	//	Jwt          string `json:"jwt"`
	//	TargetRegion string `json:"targetRegion"`
	//} `json:"mucJwtDto"`
	//MultiUserChatId       string        `json:"multiUserChatId"`
	//MultiUserChatPassword string        `json:"multiUserChatPassword"`
	//PartyId               string        `json:"partyId"`
	//PartyType             string        `json:"partyType"`
	//Restrictions          []interface{} `json:"restrictions"`
	//ScarcePositions       []interface{} `json:"scarcePositions"`
	//Warnings              []interface{} `json:"warnings"`
}

type LobbyGameConfig struct {
	AllowablePremadeSizes              []int         `json:"allowablePremadeSizes"`
	CustomLobbyName                    string        `json:"customLobbyName"`
	CustomMutatorName                  string        `json:"customMutatorName"`
	CustomRewardsDisabledReasons       []interface{} `json:"customRewardsDisabledReasons"`
	CustomSpectatorPolicy              string        `json:"customSpectatorPolicy"`
	CustomSpectators                   []interface{} `json:"customSpectators"`
	CustomTeam100                      []interface{} `json:"customTeam100"`
	CustomTeam200                      []interface{} `json:"customTeam200"`
	GameMode                           string        `json:"gameMode"`
	IsCustom                           bool          `json:"isCustom"`
	IsLobbyFull                        bool          `json:"isLobbyFull"`
	IsTeamBuilderManaged               bool          `json:"isTeamBuilderManaged"`
	MapId                              int           `json:"mapId"`
	MaxHumanPlayers                    int           `json:"maxHumanPlayers"`
	MaxLobbySize                       int           `json:"maxLobbySize"`
	MaxTeamSize                        int           `json:"maxTeamSize"`
	PickType                           string        `json:"pickType"`
	PremadeSizeAllowed                 bool          `json:"premadeSizeAllowed"`
	QueueId                            int           `json:"queueId"`
	ShouldForceScarcePositionSelection bool          `json:"shouldForceScarcePositionSelection"`
	ShowPositionSelector               bool          `json:"showPositionSelector"`
	ShowQuickPlaySlotSelection         bool          `json:"showQuickPlaySlotSelection"`
}

type GameLobby struct {
	CanStartActivity bool `json:"canStartActivity"`
	GameConfig       struct {
		AllowablePremadeSizes              []int         `json:"allowablePremadeSizes"`
		CustomLobbyName                    string        `json:"customLobbyName"`
		CustomMutatorName                  string        `json:"customMutatorName"`
		CustomRewardsDisabledReasons       []interface{} `json:"customRewardsDisabledReasons"`
		CustomSpectatorPolicy              string        `json:"customSpectatorPolicy"`
		CustomSpectators                   []interface{} `json:"customSpectators"`
		CustomTeam100                      []interface{} `json:"customTeam100"`
		CustomTeam200                      []interface{} `json:"customTeam200"`
		GameMode                           string        `json:"gameMode"`
		IsCustom                           bool          `json:"isCustom"`
		IsLobbyFull                        bool          `json:"isLobbyFull"`
		IsTeamBuilderManaged               bool          `json:"isTeamBuilderManaged"`
		MapId                              int           `json:"mapId"`
		MaxHumanPlayers                    int           `json:"maxHumanPlayers"`
		MaxLobbySize                       int           `json:"maxLobbySize"`
		MaxTeamSize                        int           `json:"maxTeamSize"`
		PickType                           string        `json:"pickType"`
		PremadeSizeAllowed                 bool          `json:"premadeSizeAllowed"`
		QueueId                            int           `json:"queueId"`
		ShouldForceScarcePositionSelection bool          `json:"shouldForceScarcePositionSelection"`
		ShowPositionSelector               bool          `json:"showPositionSelector"`
		ShowQuickPlaySlotSelection         bool          `json:"showQuickPlaySlotSelection"`
	} `json:"gameConfig"`
	Invitations []struct {
		InvitationId   string `json:"invitationId"`
		InvitationType string `json:"invitationType"`
		State          string `json:"state"`
		Timestamp      string `json:"timestamp"`
		ToSummonerId   int64  `json:"toSummonerId"`
		ToSummonerName string `json:"toSummonerName"`
	} `json:"invitations"`
	LocalMember struct {
		AllowedChangeActivity         bool          `json:"allowedChangeActivity"`
		AllowedInviteOthers           bool          `json:"allowedInviteOthers"`
		AllowedKickOthers             bool          `json:"allowedKickOthers"`
		AllowedStartActivity          bool          `json:"allowedStartActivity"`
		AllowedToggleInvite           bool          `json:"allowedToggleInvite"`
		AutoFillEligible              bool          `json:"autoFillEligible"`
		AutoFillProtectedForPromos    bool          `json:"autoFillProtectedForPromos"`
		AutoFillProtectedForRemedy    bool          `json:"autoFillProtectedForRemedy"`
		AutoFillProtectedForSoloing   bool          `json:"autoFillProtectedForSoloing"`
		AutoFillProtectedForStreaking bool          `json:"autoFillProtectedForStreaking"`
		BotChampionId                 int           `json:"botChampionId"`
		BotDifficulty                 string        `json:"botDifficulty"`
		BotId                         string        `json:"botId"`
		BotPosition                   string        `json:"botPosition"`
		BotUuid                       string        `json:"botUuid"`
		FirstPositionPreference       string        `json:"firstPositionPreference"`
		IntraSubteamPosition          interface{}   `json:"intraSubteamPosition"`
		IsBot                         bool          `json:"isBot"`
		IsLeader                      bool          `json:"isLeader"`
		IsSpectator                   bool          `json:"isSpectator"`
		PlayerSlots                   []interface{} `json:"playerSlots"`
		Puuid                         string        `json:"puuid"`
		QuickplayPlayerState          interface{}   `json:"quickplayPlayerState"`
		Ready                         bool          `json:"ready"`
		SecondPositionPreference      string        `json:"secondPositionPreference"`
		ShowGhostedBanner             bool          `json:"showGhostedBanner"`
		StrawberryMapId               interface{}   `json:"strawberryMapId"`
		SubteamIndex                  interface{}   `json:"subteamIndex"`
		SummonerIconId                int           `json:"summonerIconId"`
		SummonerId                    int64         `json:"summonerId"`
		SummonerInternalName          string        `json:"summonerInternalName"`
		SummonerLevel                 int           `json:"summonerLevel"`
		SummonerName                  string        `json:"summonerName"`
		TeamId                        int           `json:"teamId"`
		TftNPEQueueBypass             bool          `json:"tftNPEQueueBypass"`
	} `json:"localMember"`
	Members []struct {
		AllowedChangeActivity         bool          `json:"allowedChangeActivity"`
		AllowedInviteOthers           bool          `json:"allowedInviteOthers"`
		AllowedKickOthers             bool          `json:"allowedKickOthers"`
		AllowedStartActivity          bool          `json:"allowedStartActivity"`
		AllowedToggleInvite           bool          `json:"allowedToggleInvite"`
		AutoFillEligible              bool          `json:"autoFillEligible"`
		AutoFillProtectedForPromos    bool          `json:"autoFillProtectedForPromos"`
		AutoFillProtectedForRemedy    bool          `json:"autoFillProtectedForRemedy"`
		AutoFillProtectedForSoloing   bool          `json:"autoFillProtectedForSoloing"`
		AutoFillProtectedForStreaking bool          `json:"autoFillProtectedForStreaking"`
		BotChampionId                 int           `json:"botChampionId"`
		BotDifficulty                 string        `json:"botDifficulty"`
		BotId                         string        `json:"botId"`
		BotPosition                   string        `json:"botPosition"`
		BotUuid                       string        `json:"botUuid"`
		FirstPositionPreference       string        `json:"firstPositionPreference"`
		IntraSubteamPosition          interface{}   `json:"intraSubteamPosition"`
		IsBot                         bool          `json:"isBot"`
		IsLeader                      bool          `json:"isLeader"`
		IsSpectator                   bool          `json:"isSpectator"`
		PlayerSlots                   []interface{} `json:"playerSlots"`
		Puuid                         string        `json:"puuid"`
		QuickplayPlayerState          interface{}   `json:"quickplayPlayerState"`
		Ready                         bool          `json:"ready"`
		SecondPositionPreference      string        `json:"secondPositionPreference"`
		ShowGhostedBanner             bool          `json:"showGhostedBanner"`
		StrawberryMapId               interface{}   `json:"strawberryMapId"`
		SubteamIndex                  interface{}   `json:"subteamIndex"`
		SummonerIconId                int           `json:"summonerIconId"`
		SummonerId                    int64         `json:"summonerId"`
		SummonerInternalName          string        `json:"summonerInternalName"`
		SummonerLevel                 int           `json:"summonerLevel"`
		SummonerName                  string        `json:"summonerName"`
		TeamId                        int           `json:"teamId"`
		TftNPEQueueBypass             bool          `json:"tftNPEQueueBypass"`
	} `json:"members"`
	MucJwtDto struct {
		ChannelClaim string `json:"channelClaim"`
		Domain       string `json:"domain"`
		Jwt          string `json:"jwt"`
		TargetRegion string `json:"targetRegion"`
	} `json:"mucJwtDto"`
	MultiUserChatId       string        `json:"multiUserChatId"`
	MultiUserChatPassword string        `json:"multiUserChatPassword"`
	PartyId               string        `json:"partyId"`
	PartyType             string        `json:"partyType"`
	Restrictions          []interface{} `json:"restrictions"`
	ScarcePositions       []string      `json:"scarcePositions"`
	Warnings              []interface{} `json:"warnings"`
}
