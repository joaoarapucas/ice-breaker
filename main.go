package main

import (
	"github.com/gopxl/beep/speaker"
	"ice-breaker/internal/app"

	"time"
)

func main() {

	cfg := app.LoadConfig()

	trackList := app.FileNames(cfg.AudioFolder)
	app.Filter(&trackList, cfg.Blacklist, cfg.Whitelist, false)

	app.PrintFileNames(trackList)

	speaker.Init(cfg.PlaySpeed, cfg.PlaySpeed.N(time.Second/10))

	app.InitAlarm(cfg)

	app.Run(cfg, trackList)
}