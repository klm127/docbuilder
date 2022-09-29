package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filename := path.Join(pwd, "creating.txt")
	filePtr, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	filePtr.Write([]byte("hi"))
	defer filePtr.Close()
	fmt.Println("Hello World")
}
