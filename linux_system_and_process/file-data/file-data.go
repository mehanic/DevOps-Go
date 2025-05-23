package main

import (
	"os"
	"log"
)

func main() {
	// Read file to byte slice
	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}

