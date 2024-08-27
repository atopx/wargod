package opgg

import (
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const BaseUrl = "https://lol-api-champion.op.gg"

var GameModeMap = map[string]string{
	"ARAM":    "aram",
	"CLASSIC": "ranked",
}

// ListChampionSummary 英雄列表
// @param region: global/kr/jp ...
// @param mode: aram / ranked / urf ... 【大乱斗 / 排位 / 无线火力】
func ListChampionSummary(region, mode string) *ListChampionSummaryResponse {
	// https://lol-api-champion.op.gg
	api := fmt.Sprintf("%s/api/%s/champions/%s", BaseUrl, region, mode)
	resp, err := http.DefaultClient.Get(api)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var result ListChampionSummaryResponse
	if err = json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	return &result
}

// GetChampionInfo 获取英雄详情
// @param posit: top, jungle, mid, adc, support
func GetChampionInfo(id int, region, mode, posit string) *GetChampionInfoResponse {
	if mode == "aram" {
		posit = "NONE"
	}
	api := fmt.Sprintf("%s/api/%s/champions/%s/%d/%s", BaseUrl, region, mode, id, posit)
	resp, err := http.DefaultClient.Get(api)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var result GetChampionInfoResponse
	if err = json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	return &result
}

func GetAramChampionInfo(id int) *GetChampionInfoResponse {
	return GetChampionInfo(id, "global", "aram", "None")
}
