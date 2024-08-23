package core

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"time"
)

type Search struct {
	cancel chan struct{}
}

func NewSearch() *Search {
	return &Search{
		cancel: make(chan struct{}, 1),
	}
}

func (search *Search) Cancel() {
	search.cancel <- struct{}{}
}

func (search *Search) Close() {
	close(search.cancel)
}

const app = "LeagueClientUx.exe"

func (search *Search) WaitForLauncher() (*process.Process, error) {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			processes, _ := process.Processes()

			for _, p := range processes {
				name, _ := p.Name()
				if name == app {
					return p, nil
				}
			}
		case <-search.cancel:
			search.Close()
			return nil, fmt.Errorf("cancel wait")
		}
	}
}
