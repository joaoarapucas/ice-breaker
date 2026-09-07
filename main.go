package main

import (
	"fmt"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"log"
	"os"
	"time"
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

	var fileNames []string

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames
}

func main() {
	fmt.Println("hello world!")

	sfx := FileNames()
	fmt.Println("----- sounds list -----")
	for i, s := range sfx {
		fmt.Printf("%d - %s\n", i+1, s)
	}

	file, err := os.Open("sfx/" + sfx[0])
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	sr := format.SampleRate * 1 //audio speed
	speaker.Init(sr, sr.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
