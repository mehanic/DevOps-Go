package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Check for correct number of arguments
	if len(os.Args) != 2 {
		fmt.Printf("%s basepath\n", os.Args[0])
		fmt.Println()
		os.Exit(1)
	}

	path := os.Args[1]

	// Map to store file type counts
	statArray := make(map[string]int)

	// Walk through the directory and process each file
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a file
		if !info.IsDir() {
			// Run the "file" command to get the file type
			out, err := exec.Command("file", "-b", path).Output()
			if err != nil {
				log.Printf("Failed to get file type for %s: %v", path, err)
				return err
			}

			// Clean the file type output
			fileType := strings.TrimSpace(string(out))

			// Increment the count for the file type
			statArray[fileType]++
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", path, err)
	}

	// Print the results
	for ftype, count := range statArray {
		fmt.Printf("%s: %d\n", ftype, count)
	}
}

