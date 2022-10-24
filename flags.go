package main

import (
	"flag"
	"log"
)

func ParseFlags() {
	out := flag.String("out", "", "A relative path to the output directory.")
	from := flag.String("from", "", "A relative path to the input directory.")
	logs := flag.String("log", "", "A relative path the logfile directory.")
	create := flag.Bool("make", false, "Creates missing reference and content files.")
	flag.Parse()
	if *out != "" {
		Settings.Outpath = *out
		log.Println("Set outpath to", *out, "from user flag.")
	}
	if *from != "" {
		Settings.Inpath = *from
		log.Println("Set inpath to", *from, "from user flag.")
	}
	if *logs != "" {
		Settings.LogPath = *logs
	}
	if *create == true {
		Settings.CreateMissingFiles = true
		log.Println("Set CreateMissingFiles to true.")
	}
}
