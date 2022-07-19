package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	lib "spotify-playlist-parser/lib"
	"strings"
)

const (
	filename = "Playlist1.json"
	dirname  = "My playlists"
)

func main() {
	//open file
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Printf("Error reading file %s: %s", filename, err)
	}
	defer jsonFile.Close()

	//read all strings
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var file lib.File
	json.Unmarshal(byteValue, &file)

	//create directory
	if err := os.Mkdir(dirname, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for _, playlist := range file.Playlists {
		//check if any "/" in names
		playlist.Name = strings.Replace(playlist.Name, "/", "-", -1)

		filename := dirname + "/" + playlist.Name + ".txt"

		//generate file
		log.Printf("Generating file: \"%s\"", filename)
		playlistFile, err := os.Create(filename)
		if err != nil {
			log.Printf("Error create file %s: %s", filename, err)
		}
		defer playlistFile.Close()

		//fill file with songs
		for _, track := range playlist.Items {
			if track.LocalTrack == nil {
				playlistFile.WriteString(track.Track.ArtistName + "-" + track.Track.TrackName + "\n")
			} else {
				log.Printf("Local file found: %s", track.LocalTrack.URI)
			}

		}
	}

}
