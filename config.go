package main

import (
	"io/fs"
	"log"
	"os"
	"path"
)

// Config holds configuration options for docbuilder.
type Config struct {
	useDefaults       bool
	Outpath           string
	Outname           string
	FolderNames       []string
	ContentFileName   string
	ReferenceFileName string
	TermSections      string
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
	fullPath := path.Join(pwd, c.Outpath)
	valid := fs.ValidPath(fullPath)
	if !valid {
		l("Bad output path: " + c.Outpath)
		c.Outpath = "."
		l("Defaulting to current directory.")
	}
	if c.Outname == "" {
		l("Bad outname, empty.")
		c.Outname = "out"
		l("Defaulting to 'out'")
	}
	var validatedFolders []string
	for _, s := range c.FolderNames {
		fullPath = path.Join(pwd, s)
		valid := fs.ValidPath(fullPath)
		if !valid {
			l("Bad output path: ", fullPath)
		}
		validatedFolders = append(validatedFolders, s)
	}
	c.FolderNames = validatedFolders
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
