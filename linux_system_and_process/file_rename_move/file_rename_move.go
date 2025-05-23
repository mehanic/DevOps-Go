package main

import (
	"log"
	"os"
)

func main() {

	originalPath := "demo.txt"
	newPath := "test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

