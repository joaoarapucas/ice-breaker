package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// gets all sounds file names
func FileNames(folderPath string) []string {
	files, err := os.ReadDir(folderPath)
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

func Filter(trackList *[]string, blacklist []string, whitelist []string, exclusive bool) {
	/*
		blacklisted sounds will not be played
		whitelisted sounds will be the only ones to be played
	*/

	if trackList == nil || len(*trackList) == 0 {
		return
	}

	// has whitelist
	if len(whitelist) > 0 {
		// creates whitelist hash map
		whiteMap := make(map[string]bool, len(whitelist))
		for _, item := range whitelist {
			whiteMap[item] = true
		}

		n := 0
		//if track is whitelisted, add to trackList
		for _, track := range *trackList {
			if whiteMap[track] {
				(*trackList)[n] = track
				n++
			}
		}
		//resizes trackList to new whitelist size
		*trackList = (*trackList)[:n]

		if exclusive {
			return
		}
	}

	// has blacklist
	if len(blacklist) > 0 {
		// creates blacklist hash map
		blackMap := make(map[string]bool, len(blacklist))
		for _, item := range blacklist {
			blackMap[item] = true
		}

		n := 0
		//if track is not blacklisted, add to trackList
		for _, track := range *trackList {
			if !blackMap[track] {
				(*trackList)[n] = track
				n++
			}
		}
		//resizes trackList to new blacklist size
		*trackList = (*trackList)[:n]
	}
}

func PrintFileNames(trackList []string) {
	fmt.Println("----- sounds list -----")
	for i, s := range trackList {
		fmt.Printf("%d - %s\n", i+1, s)
	}
	if len(trackList) == 0 {
		fmt.Printf("no sound found...")
	}
	fmt.Println("-----------------------")
}
