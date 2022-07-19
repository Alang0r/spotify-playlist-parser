package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	lib "spotify-playlist-parser/lib"
)

const (
	filename = "playlist1.json"
)

func main() {
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Printf("Error reading file %s: %s", filename, err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var file lib.File
	json.Unmarshal(byteValue, file)
	


}
