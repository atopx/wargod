package main

import (
	"fmt"
	"wargod/core"
)

func main() {
	p, err := core.NewSearch().WaitForLauncher()
	if err != nil {
		panic(err)
	}
	auth, err := core.NewGameAuth(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("PID:", auth.Pid)
	fmt.Printf("https://riot:%s@127.0.0.1:%s\n", auth.Token, auth.Port)
}
