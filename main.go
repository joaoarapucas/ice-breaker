package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"log"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ----- DEFAULT CONFIG ----- //
var playSpeed = beep.SampleRate(44100) * 1
var audioFolder = "./sfx/"
var randomChance = /* 1 in */ 60
var tickSpeed = time.Second * 1
var alarmHour = 17
var alarmMinute = 0

type Config struct {
	AudioFolder  string `toml:"audio_folder"`
	RandomChance int    `toml:"random_chance"`
	TickSpeed    int    `toml:"tick_speed"`
	PlaySpeed    int    `toml:"play_speed"`
	AlarmHour    int    `toml:"alarm_hour"`
	AlarmMinute  int    `toml:"alarm_minute"`
}

func LoadConfig() Config {
	var cfg Config
	if _, err := toml.DecodeFile("./config.toml", &cfg); err != nil {
		log.Fatalf("error when loading config: %v", err)
	}
	return cfg
}

// gets all sound effects file names
func FileNames() []string {
	files, err := os.ReadDir(audioFolder)
	if err != nil {
		log.Fatal(err)
	}

	//valid audio extensions
	audioExts := map[string]bool{
		".mp3":  true,
		".wav":  true,
		".ogg":  true,
		".flac": true,
	}

	var fileNames []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		//if is audio file
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if audioExts[ext] {
			fileNames = append(fileNames, file.Name())
		} else {
			log.Printf("file %s skipped - not an audio file", file.Name())
		}
	}
	return fileNames
}

func PlaySound(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	streamer, _, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
		file.Close()
		return
	}

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		streamer.Close()
		file.Close()
	})))

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

func main() {
	fmt.Println("hello world!")

	cfg := LoadConfig()
	playSpeed = beep.SampleRate(cfg.PlaySpeed)
	audioFolder = cfg.AudioFolder
	randomChance = cfg.RandomChance
	tickSpeed = time.Second * time.Duration(cfg.TickSpeed)
	alarmHour = cfg.AlarmHour
	alarmMinute = cfg.AlarmMinute

	sounds := FileNames()
	fmt.Println("----- sounds list -----")
	for i, s := range sounds {
		fmt.Printf("%d - %s\n", i+1, s)
	}
	if len(sounds) == 0 {
		fmt.Printf("no sound found...")
	}
	fmt.Println("-----------------------")

	speaker.Init(playSpeed, playSpeed.N(time.Second/10))

	if !(alarmHour < 0 || alarmMinute < 0) {
		go ScheduledAlarm(alarmHour, alarmMinute, audioFolder, "telephone_ring.mp3")
	}

	timer := time.NewTicker(tickSpeed) //trigger every tick
	defer timer.Stop()
	i := 1

	//core loop
	for range timer.C {
		fmt.Print("tick ", i)
		i++
		for _, s := range sounds {
			if rand.IntN(randomChance) == 0 {
				go func() {
					PlaySound(audioFolder + s)
				}()
				fmt.Print(" - played " + s)
			}
		}
		fmt.Print("\n")
	}
}
