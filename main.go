package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	lib "spotify-playlist-parser/lib"
)

const (
	filename = "Playlist1.json"
	dirname = "My playlists"
)

func main() {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Printf("Error reading file %s: %s", filename, err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var file lib.File
	json.Unmarshal(byteValue, &file)

	if err := os.Mkdir(dirname, os.ModePerm); err != nil {
        log.Fatal(err)
    }

	for _, playlist := range file.Playlists {
		filename := dirname + "/" + playlist.Name + ".txt"
		log.Printf("Generating file: \"%s\"", filename)
		playlistFile, err := os.Create(filename)
		if err != nil {
			log.Printf("Error create file %s: %s", filename, err)
		}
		defer playlistFile.Close()

		for _, track := range playlist.Items {
			if track.LocalTrack == nil {
				playlistFile.WriteString(track.Track.ArtistName + "-" + track.Track.TrackName + "\n")
			} else {
				log.Printf("Local file found: %s", track.LocalTrack.URI)
			}

		}
	}

}
