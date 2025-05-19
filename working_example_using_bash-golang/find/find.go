package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	startDir := "/usr/local/bin"
	outputFile := "find.sh"

	// Create or truncate the output file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", outputFile, err)
		return
	}
	defer file.Close()

	// Variable to track if any files are found
	foundFiles := false

	// Walk through the directory
	err = filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Debug: Print current file/directory being processed
		fmt.Printf("Processing: %s\n", path)

		// Skip the .git directory and find.sh file
		if info.IsDir() && (info.Name() == ".git" || info.Name() == "find.sh") {
			return filepath.SkipDir
		}

		// Check if the file matches the pattern b*.sh
		if !info.IsDir() && strings.HasPrefix(info.Name(), "b") && strings.HasSuffix(info.Name(), ".sh") {
			foundFiles = true // Set the flag to true
			// Write the matched file path to the output file
			if _, err := file.WriteString(path + "\n"); err != nil {
				return fmt.Errorf("error writing to file %s: %v", outputFile, err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through the directory: %v\n", err)
		return
	}

	// Write an empty line at the end of the output file if any files were found
	if foundFiles {
		if _, err := file.WriteString("\n"); err != nil {
			fmt.Printf("Error writing final newline to %s: %v\n", outputFile, err)
		}
	} else {
		fmt.Println("No files matching the criteria were found.")
	}

	fmt.Printf("Filtered paths written to %s\n", outputFile)
}

