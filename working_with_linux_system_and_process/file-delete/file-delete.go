package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}



func main1() {

	err := os.Remove("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
}
