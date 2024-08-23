package opgg

import (
	"fmt"
	"testing"
)

func TestListChampion(t *testing.T) {
	result := ListChampionSummary("global", "aram")
	fmt.Printf("大乱斗: %+v\n", *result)
	result = ListChampionSummary("global", "ranked")
	fmt.Printf("排位: %+v\n", *result)
}

func TestGetChampionInfo(t *testing.T) {
	result := GetChampionInfo(412, "global", "aram", "none")

	fmt.Printf("大乱斗: %+v\n", *result)
	result = GetChampionInfo(412, "global", "ranked", "support")
	fmt.Printf("排位: %+v\n", *result)
}
