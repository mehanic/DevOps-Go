package main

import (
	"log"
	"os"
)

func main() {
	if err := os.MkdirAll("a/b/c/d", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

