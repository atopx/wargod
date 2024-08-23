package model

import "github.com/goccy/go-json"

type Rune struct {
	Id              int    `json:"id,omitempty"`
	Name            string `json:"name"`
	IsActive        bool   `json:"isActive"`
	Current         bool   `json:"current"`
	Order           int    `json:"order"`
	PrimaryStyleId  int    `json:"primaryStyleId"`
	SubStyleId      int    `json:"subStyleId"`
	SelectedPerkIds []int  `json:"selectedPerkIds"` // 4+2+3 => 主4+副2+通用3
}

func NewRune(data []byte) *Rune {
	r := new(Rune)
	_ = json.Unmarshal(data, &r)
	return r
}

func (r *Rune) MustMarshal() []byte {
	data, _ := json.Marshal(r)
	return data
}

// 下面都是原始结构, 用不上
//type Rune struct {
//	AutoModifiedSelections   []interface{}    `json:"autoModifiedSelections,omitempty"` // 不知道啥玩意
//	Current                  bool             `json:"current,omitempty"`
//	Id                       int              `json:"id,omitempty"`
//	IsActive                 bool             `json:"isActive,omitempty"`
//	IsDeletable              bool             `json:"isDeletable,omitempty"`
//	IsEditable               bool             `json:"isEditable,omitempty"`
//	IsRecommendationOverride bool             `json:"isRecommendationOverride,omitempty"`
//	IsTemporary              bool             `json:"isTemporary,omitempty"`
//	IsValid                  bool             `json:"isValid,omitempty"`
//	LastModified             int64            `json:"lastModified,omitempty"`
//	Name                     string           `json:"name,omitempty"`
//	Order                    int              `json:"order,omitempty"`
//	PrimaryStyleIconPath     string           `json:"primaryStyleIconPath,omitempty"`
//	PrimaryStyleId           int              `json:"primaryStyleId,omitempty"`
//	PrimaryStyleName         string           `json:"primaryStyleName,omitempty"`
//	QuickPlayChampionIds     []interface{}    `json:"quickPlayChampionIds,omitempty"` // 不知道啥玩意
//	RecommendationChampionId int              `json:"recommendationChampionId,omitempty"`
//	RecommendationIndex      int              `json:"recommendationIndex,omitempty"`
//	RuneRecommendationId     string           `json:"runeRecommendationId,omitempty"`
//	SecondaryStyleIconPath   string           `json:"secondaryStyleIconPath,omitempty"`
//	SecondaryStyleName       string           `json:"secondaryStyleName,omitempty"`
//	SelectedPerkIds          []int           `json:"selectedPerkIds,omitempty"` // 4+2+3 => 主4+副2+通用3
//	SubStyleId               int              `json:"subStyleId,omitempty"`
//	TooltipBgPath            string           `json:"tooltipBgPath,omitempty"`
//	UiPerks                  []RuneUiPerk     `json:"uiPerks,omitempty"`
//	PageKeystone             RunePageKeystone `json:"pageKeystone,omitempty"`
//}

//type RunePageKeystone struct {
//	IconPath string `json:"iconPath,omitempty"`
//	Id       int    `json:"id,omitempty"`
//	Name     string `json:"name,omitempty"`
//	SlotType string `json:"slotType,omitempty"`
//	StyleId  int    `json:"styleId,omitempty"`
//}
//
//type RuneUiPerk struct {
//	IconPath string `json:"iconPath,omitempty"`
//	Id       int    `json:"id,omitempty"`
//	Name     string `json:"name,omitempty"`
//	SlotType string `json:"slotType,omitempty"`
//	StyleId  int    `json:"styleId,omitempty"`
//}
