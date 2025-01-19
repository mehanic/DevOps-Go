package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Mkdir("a", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

