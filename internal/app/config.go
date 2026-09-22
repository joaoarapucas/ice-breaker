package app

import (
	"github.com/BurntSushi/toml"
	"github.com/gopxl/beep"
	"log"
	"time"
)

type Config struct {
	AudioFolder  string          `toml:"audio_folder"`
	RandomChance int             `toml:"random_chance"`
	TickSpeed    time.Duration   `toml:"tick_speed"`
	PlaySpeed    beep.SampleRate `toml:"play_speed"`

	AlarmHour   int    `toml:"alarm_hour"`
	AlarmMinute int    `toml:"alarm_minute"`
	AlarmSound  string `toml:"alarm_sound"`

	Blacklist []string `toml:"blacklist"`
	Whitelist []string `toml:"whitelist"`
}

func LoadConfig() Config {
	var cfg Config
	if _, err := toml.DecodeFile("./config.toml", &cfg); err != nil {
		log.Fatalf("error when loading config: %v", err)
	}

	cfg.TickSpeed = cfg.TickSpeed * time.Second

	return cfg
}
