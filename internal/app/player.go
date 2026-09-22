package app

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
	"log"
	"os"
	// "path/filepath"
)

// func Decode(path string) {

// 	streamer, err: beep.Streamer

// 	file, err := os.Open(path)

// 	ext := strings.ToLower(filepath.Ext(path))

// 		".wav":  true,
// 		".ogg":  true,
// 		".flac": true,
// 	siwtch ext {
// 		case ".mp3":
// 			streamer, _, err := mp3.Decode(file)
// 		case ".ogg":
// 			streamer, _, err :=
// 		case ".flac":

// 	}
// }

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
