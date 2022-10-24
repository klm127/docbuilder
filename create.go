package main

import (
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
			log.Println("Failed to create " + inpath)
		}
	}
	for _, folder := range Settings.FolderNames {
		folderPath := path.Join(inpath, folder)
		_, err = os.Stat(folderPath)
		if err != nil {
			err = os.Mkdir(folderPath, 0755)
			if err != nil {
				log.Println("Failed to create" + folderPath)
			} else {
				if Settings.CreateMissingFiles {
					createDirectoryFiles(folderPath)
				}
			}
		} else {
			if Settings.CreateMissingFiles {
				createDirectoryFiles(folderPath)
			}
		}
	}
}

// createDirectoryFiles is an internal function to create the inner files
func createDirectoryFiles(folderPath string) {
	contPath := path.Join(folderPath, Settings.ContentFileName)
	_, err := os.Stat(contPath)
	if err != nil {
		file, err := os.Create(contPath)
		if err != nil {
			log.Println("Failed to create " + contPath)
		} else {
			file.Close()
		}
	}
	refPath := path.Join(folderPath, Settings.ReferenceFileName)
	_, err = os.Stat(refPath)
	if err != nil {
		file, err := os.Create(refPath)
		if err != nil {
			log.Println("Failed to create " + refPath)
		} else {
			file.Close()
		}
	}
}
