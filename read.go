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

// ReadSections calls readSection for each folder specified by Settings. It returns the group as an array of SectionReads.
func ReadSections() []SectionRead {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get pwd in ReadSections")
	}
	var readresults []SectionRead
	for _, name := range Settings.FolderNames {
		fpath := path.Join(pwd, name)
		sec := readSection(fpath)
		sec.Name = name
		readresults = append(readresults, sec)
	}
	return readresults
}

// readSection reads one of the input folders. It generates a raw read result called SectionRead containing the folder name, raw content, and raw references data. That data must be further parsed by doc builder.
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
		readResult.Content = ""
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
