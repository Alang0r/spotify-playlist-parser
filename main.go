package main

import (
	"encoding/json"
	"log"
	"os"
	lib "spotify-playlist-parser/lib"
)


const(
	filename = "playlist1.json"
)
func main() {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error reading file %s: %s", filename, err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var playlists []lib.Playlist
	json.Unmarshal(byteValue, playlists)


	

}
