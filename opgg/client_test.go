package opgg

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/atopx/clever"
)

func TestListChampion(t *testing.T) {
	result := ListChampionSummary("global", "aram")
	fmt.Printf("大乱斗: %+v\n", *result)
	result = ListChampionSummary("global", "ranked")
	fmt.Printf("排位: %+v\n", *result)
}

func TestGetChampionInfo(t *testing.T) {
	result := GetChampionInfo(201, "global", "aram", "none")
	data, _ := json.Marshal(&result)
	fmt.Println(clever.String(data))
}
