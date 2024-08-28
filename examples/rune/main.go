package main

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"time"
	"wargod/core"
	"wargod/model"
)

// 测试设置符文页
func main() {
	ga := core.NewGameAuthWithToken("2343", "nSG8Dnix2edABLtIl4Jqdg")
	ctx := context.Background()
	client, err := core.New(ctx, time.Second, ga)
	if err != nil {
		log.Fatal(err)
	}
	// 获取当前符文页
	data, _, err := client.Get("/lol-summoner/v1/current-summoner")
	if err != nil {
		log.Fatalf("wait for game launcher error: %s", err)
	}
	summoner := new(model.Summoner)
	_ = json.Unmarshal(data, summoner)
	fmt.Println(summoner)
}
