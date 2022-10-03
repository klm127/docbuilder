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
	readSettings()
	CreateDirectories()
}

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
