package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

// CreateDirectoriesIfNotExisting checks the Setting for the list of directories that will be built from. If those directories do not exist, it creates them. It also creates the content file and reference file for each directory.
func CreateDirectoriesIfNotExisting() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	inpath := path.Join(pwd, Settings.Inpath)
	_, err = os.Stat(inpath)
	if err != nil {
		err = os.Mkdir(inpath, 0755)
		if err != nil {
			fmt.Println("Failed to create " + inpath)
		}
	}
	for _, folder := range Settings.FolderNames {
		folderPath := path.Join(inpath, folder)
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
