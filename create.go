package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func CreateDirectoriesIfNotExisting() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for _, folder := range Settings.FolderNames {
		folderPath := path.Join(pwd, folder)
		_, err := os.Stat(folderPath)
		if err != nil {
			err = os.Mkdir(folderPath, 0755)
			if err != nil {
				fmt.Println("Failed to create" + folderPath)
			} else {
				contPath := path.Join(folderPath, Settings.ContentFileName)
				refPath := path.Join(folderPath, Settings.ReferenceFileName)
				file, err := os.Create(contPath)
				if err != nil {
					fmt.Println("Failed to create " + contPath)
				} else {
					file.Close()
				}
				file, err = os.Create(refPath)
				if err != nil {
					fmt.Println("Failed to create " + refPath)
				} else {
					file.Close()
				}
			}
		}
	}
}
