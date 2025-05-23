package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Define the file path
	filePath := "/tmp/random_file.txt"

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist, stopping the script!")
		os.Exit(1)
	}

	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Print the file content
	fmt.Print(string(content))
}

