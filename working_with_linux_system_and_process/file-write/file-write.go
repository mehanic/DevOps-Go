package main

import (
	"os"
	"log"
)

func main() {
	err := os.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

