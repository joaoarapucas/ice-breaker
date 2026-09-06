package main

import (
	"fmt"
	//	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Print("hello world!")
	file, err := os.Open("sfx.mp3")
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

	speaker.Play(streamer)

	select {}
}
