package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

type SectionRead struct {
	Name       string
	Content    string
	References string
}

func ReadSections() []SectionRead {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get pwd in ReadSections")
	}
	var readresults []SectionRead
	for _, name := range Settings.FolderNames {
		fpath := path.Join(pwd, name)
		sec := readSection(fpath)
		readresults = append(readresults, sec)
	}
	return readresults
}

func readSection(folderpath string) SectionRead {
	var readResult SectionRead
	contentPath := path.Join(folderpath, Settings.ContentFileName)
	refPath := path.Join(folderpath, Settings.ReferenceFileName)
	bytes, err := os.ReadFile(contentPath)
	if err == nil {
		parsed := string(bytes)
		readResult.Content = parsed
	} else {
		fmt.Println("Error reading " + contentPath)
	}
	bytes, err = os.ReadFile(refPath)
	if err == nil {
		parsed := string(bytes)
		readResult.References = parsed
	} else {
		fmt.Println("Error reading " + refPath)
	}
	return readResult
}
