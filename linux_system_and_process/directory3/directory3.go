package main

import (
	"fmt"
	"os"
)

func makeLogsDir() {
	err := os.Mkdir("logs", 0755)
	if err != nil {
		fmt.Println("Logs directory already exists")
	}
}

func makeOutputDir() {
	err := os.Mkdir("output", 0755)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// makeLogsDir()
	makeOutputDir()
}
