package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("err happened:", err)
		// fatal calls os.Exit(1) after writing the log message
	}
}

// 2025/02/14 15:08:12 err happened: open no-file.txt: no such file or directory
// exit status 1

