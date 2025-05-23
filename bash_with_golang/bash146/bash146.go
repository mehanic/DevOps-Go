package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func fileCount() int {
	// Read the contents of the current directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	// Initialize a counter for regular files
	fileCounter := 0

	// Loop through the entries in the directory
	for _, file := range files {
		// Only count regular files
		if !file.IsDir() {
			fileCounter++
		}
	}

	// Return the count of regular files
	return fileCounter
}

func main() {
	numberOfFiles := fileCount() // Call the fileCount function
	fmt.Println(numberOfFiles)    // Print the number of regular files
}

