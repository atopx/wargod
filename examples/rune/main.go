package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"wargod/core"
	"wargod/model"
	"wargod/opgg"
)

// 测试设置符文页
func main() {
	ga := core.NewGameAuthWithToken("5261", "YJB50cyMXm-an903Xy-dIA")
	ctx := context.Background()
	client, err := core.New(ctx, time.Second, ga)
	if err != nil {
		log.Fatal(err)
	}
	// 获取当前符文页
	data, err := client.Get("/lol-perks/v1/currentpage")
	if err != nil {
		log.Fatal(err)
	}
	r := model.NewRune(data)
	fmt.Println("当前符文页:", r.Id, r.Order, r.PrimaryStyleId, r.SubStyleId, r.SelectedPerkIds)

	// 获取OPGG符文 锤石=412
	info := opgg.GetAramChampionInfo(412)

	// 根据算法获取最佳符文
	opggBestRune := info.Data.BestRune()
	fmt.Printf("%+v\n", opggBestRune)

	// 设置符文页
	r.Name = "atopx"
	r.Order = 0
	r.PrimaryStyleId = opggBestRune.PrimaryPageId
	r.SubStyleId = opggBestRune.SecondaryPageId
	r.Current = true
	r.IsActive = true
	r.SelectedPerkIds = make([]int, 0, 9)
	r.SelectedPerkIds = append(r.SelectedPerkIds, opggBestRune.PrimaryRuneIds...)
	r.SelectedPerkIds = append(r.SelectedPerkIds, opggBestRune.SecondaryRuneIds...)
	r.SelectedPerkIds = append(r.SelectedPerkIds, opggBestRune.StatModIds...)
	fmt.Printf("%+v\n", r)
	resp, err := client.Put(fmt.Sprintf("/lol-perks/v1/pages/%d", r.Id), r.MustMarshal())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
}
