package app

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func Run(cfg Config, trackList []string) {
	timer := time.NewTicker(cfg.TickSpeed)
	defer timer.Stop()

	i := 1
	for range timer.C {
		fmt.Print("tick ", i)
		i++
		playRandomTracks(cfg, trackList)
		fmt.Print("\n")
	}
}

func playRandomTracks(cfg Config, trackList []string) {
	for _, s := range trackList {
		if rand.IntN(cfg.RandomChance) == 0 {
			go PlaySound(cfg.AudioFolder + s)
			fmt.Print(" - played " + s)
		}
	}
}
