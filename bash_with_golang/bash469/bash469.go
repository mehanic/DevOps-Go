package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if a file name argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run empty-file.go <file-name>")
		return
	}

	// Grab the first argument
	fileName := os.Args[1]

	// Check if the file exists
	if _, err := os.Stat(fileName); err == nil {
		// File exists; truncate it to empty
		if err := os.Truncate(fileName, 0); err != nil {
			fmt.Printf("Something went wrong, please check %s!\n", fileName)
			os.Exit(1)
		}
	} else {
		// File does not exist; create it
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Something went wrong, please check %s!\n", fileName)
			os.Exit(1)
		}
		defer file.Close()
	}

	fmt.Printf("Success, file %s is now empty.\n", fileName)
}

