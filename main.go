package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

var Settings = NewConfig()

func main() {
	SetFileLogging()
	log.Println("~ New Docbuilder session. ~")
	readSettings()
	ParseFlags()
	CreateDirectoriesIfNotExisting()
	_ = ReadSections()
	SaveLog()
}

// readSettings reads the dobuilder.json file in the process directory, looking for a docbuilder.json file. It configures docbuilder based on the settings in that file.
func readSettings() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filename := path.Join(pwd, "docbuilder.json")
	fmt.Println(filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Ensure %s exists.\n%v", filename, err)
	}
	if err != nil {
		log.Fatalf("Failed to read .json. %v", err)
	}
	err = json.Unmarshal(content, &Settings)
	if err != nil {
		log.Fatalf("Failed to read .json. %v", err)
	}
	Settings.Validate()
}
