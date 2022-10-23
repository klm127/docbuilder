package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
)

type TLogger struct {
	buffer   bytes.Buffer
	openFile *os.File
}

func SaveLog() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to save log; could not get process working directory!")
		log.Fatal(err)
	} else {
		path := path.Join(pwd, Settings.LogPath, "logs.log")
		fmt.Println("Saving logs to " + path)
		file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		file.Write(logger.buffer.Bytes())
		defer file.Close()
	}
}

var logger TLogger

func SetFileLogging() {
	log.SetOutput(&logger.buffer)
}
