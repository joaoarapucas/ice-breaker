package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

/*
TODO - JSON
blacklist
whitelist
sounds folder
if key pressed plays sound
global random time ratio
*/

// gets all sound effects file names
func FileNames() []string {
	files, err := os.ReadDir("sfx")
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

	var s beep.Streamer = streamer

	speaker.Play(beep.Seq(s, beep.Callback(func() {
		streamer.Close()
		file.Close()
	})))

}

func main() {
	fmt.Println("hello world!")

	sfx := FileNames()
	fmt.Println("----- sounds list -----")
	for i, s := range sfx {
		fmt.Printf("%d - %s\n", i+1, s)
	}

	//play speed
	const sampleRate = beep.SampleRate(44100)

	speaker.Init(sampleRate*2, sampleRate.N(time.Second/10))

	PlaySound("sfx/" + sfx[2])
	PlaySound("sfx/" + sfx[3])
	select {}
}
