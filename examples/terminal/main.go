package main

import (
	"context"
	"wargod/game"
)

func main() {
	ctx := context.Background()
	lol := game.New(ctx)
	lol.Start()
}
