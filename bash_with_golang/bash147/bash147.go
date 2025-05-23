package main

import (
	"fmt"
	"io/ioutil"
	"log"
//	"os"
)

// fileCount counts the number of regular files in the specified directory
func fileCount(dir string) {
	// Read the contents of the specified directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory %s: %v", dir, err)
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

	// Print the directory and the count of regular files
	fmt.Printf("%s:\n%d\n", dir, fileCounter)
}

func main() {
	// Call fileCount for each specified directory
	fileCount("/etc")
	fileCount("/var")
	fileCount("/usr/bin")
}

