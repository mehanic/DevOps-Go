package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	// Define the directory and file pattern
	dirPath := "/path/to/dir/"
	filePattern := "*.sql"

	// Use filepath.Glob to get all .sql files in the directory
	files, err := filepath.Glob(filepath.Join(dirPath, filePattern))
	if err != nil {
		log.Fatalf("Failed to get files: %v", err)
	}

	// Process each file
	for _, file := range files {
		// Code block for processing each file
		processFile(file)
	}
}

// Example function to process each file
func processFile(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Error reading file %s: %v", file, err)
		return
	}

	// Here you can add your logic to process the content
	fmt.Printf("Processing file: %s\n", file)
	// For example, just printing the content
	fmt.Println(string(content))
}

