package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Check for command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run if_then_else_proper.go <file-name>")
		os.Exit(1)
	}

	// Grab the file name from command line arguments
	fileName := os.Args[1]

	// Check if the file exists
	if _, err := os.Stat(fileName); err == nil {
		// If the file exists, read and print its content
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
		fmt.Print(string(content)) // Print the file content
	} else {
		// If the file does not exist
		fmt.Println("File does not exist, stopping the script!")
		os.Exit(1)
	}
}

