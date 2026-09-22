package app

import (
	"fmt"
	"time"
)

func InitAlarm(cfg Config) {
	if cfg.AlarmHour < 0 || cfg.AlarmMinute < 0 {
		fmt.Println("alarm will not be played.")
	} else {
		go ScheduledAlarm(cfg.AlarmHour, cfg.AlarmMinute, cfg.AudioFolder, cfg.AlarmSound)
		fmt.Printf("alarm scheduled to %d: %d\n", cfg.AlarmHour, cfg.AlarmMinute)
	}
}

func ScheduledAlarm(hour int, minute int, dir string, file string) {
	timer := time.NewTicker(time.Second) //trigger every second
	defer timer.Stop()

	for range timer.C {
		if time.Now().Hour() == hour && time.Now().Minute() == minute {
			fmt.Println("bateu !!!")
			PlaySound(dir + file)
			return
		}
	}
}
