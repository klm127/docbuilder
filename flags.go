package main

import (
	"flag"
)

func ParseFlags() {
	n := flag.String("testholder", "11", "nm")
	flag.Parse()
	_ = n
	//log.Printf("Log value was %s\n", *n)
}
