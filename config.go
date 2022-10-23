package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
)

// Config holds configuration options for docbuilder.
type Config struct {
	LogPath           string
	useDefaults       bool
	Outpath           string
	Outname           string
	Inpath            string
	FolderNames       []string
	ContentFileName   string
	ReferenceFileName string
	TermSections      []string
	OutputType        string
}

// NewConfig returns a new Config object.
func NewConfig() Config {
	var c Config
	c.useDefaults = true
	return c
}

// Validate checks all values in Config to ensure they will work the program. If Config.useDefaults is True, default values will replace invalid values.
func (c *Config) Validate() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting pwd: %v", err)
	}
	var l func(...any)
	if c.useDefaults {
		l = log.Println
	} else {
		l = log.Fatal
	}
	if c.LogPath == "" {
		l("Invalid logpath: " + c.LogPath)
		l("Defaulting to current directory.")
		c.LogPath = "."
	}
	_, err = os.Stat(c.LogPath)
	if err != nil {
		l("Logpath directory " + c.LogPath + " doesn't exist.")
		l("Creating that directory.")
		err = os.Mkdir(c.LogPath, os.ModeDir)
		if err != nil {
			fmt.Println("couldn't make log directory of " + err.Error())

		}
	}
	outputPath := path.Join(pwd, c.Outpath)
	isValid := fs.ValidPath(outputPath)
	if !isValid {
		l("Bad output path: " + c.Outpath)
		c.Outpath = "."
		l("Defaulting to current directory.")
	}
	if c.Outname == "" {
		l("Bad outname, empty.")
		c.Outname = "out"
		l("Defaulting to 'out'")
	}
	if c.Inpath == "" {
		l("Bad input path, empty.")
		c.Inpath = "."
		l("Defaulting to '.', current directory.")
	}
	inputPath := path.Join(pwd, c.Inpath)
	var validatedInputFolders []string
	for _, s := range c.FolderNames {
		outputPath = path.Join(inputPath, s)
		valid := fs.ValidPath(outputPath)
		if !valid {
			l("Bad input path: ", outputPath)
		}
		validatedInputFolders = append(validatedInputFolders, s)
	}
	c.FolderNames = validatedInputFolders
	if c.ReferenceFileName == "" {
		l("No reference file name provided.")
		c.ReferenceFileName = "references.txt"
		l("Defaulting to 'references.txt'")
	}
	if c.ContentFileName == "" {
		l("No content file name provided.")
		c.ContentFileName = "content.md"
		l("Defaulting to 'content.md'")
	}
	if c.OutputType != "md" && c.OutputType != "tex" && c.OutputType != "txt" {
		l("No output type provided.")
		c.OutputType = "md"
		l("Defaulting to 'md'")
	}
}
