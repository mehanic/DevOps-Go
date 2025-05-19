package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// isFile checks if a given path is a file and exists.
func isFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("%s does not seem to be a file\n", filename)
		os.Exit(2)
	}
}

// countLines returns the number of lines in the given file.
func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount, scanner.Err()
}

// cleanFile removes empty lines and lines starting with #, then counts lines before and after.
func cleanFile(filename string) {
	isFile(filename)

	// Count lines before cleaning
	before, err := countLines(filename)
	if err != nil {
		fmt.Printf("Error counting lines: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("The file %s starts with %d lines\n", filename, before)

	// Open the file for reading and create a backup file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a backup file
	backupFile, err := os.Create(filename + ".bak")
	if err != nil {
		fmt.Printf("Error creating backup file: %v\n", err)
		os.Exit(1)
	}
	defer backupFile.Close()

	// Open the original file for writing the cleaned content
	tempFile, err := os.CreateTemp("", "cleaned_file")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		os.Exit(1)
	}
	defer tempFile.Close()

	// Process the file to remove empty and comment lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Remove lines that are empty or start with "#"
		if strings.TrimSpace(line) == "" || strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}
		// Write valid lines to the temp file and the backup file
		fmt.Fprintln(tempFile, line)
		fmt.Fprintln(backupFile, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
		os.Exit(1)
	}

	// Replace the original file with the cleaned temp file
	err = os.Rename(tempFile.Name(), filename)
	if err != nil {
		fmt.Printf("Error replacing original file: %v\n", err)
		os.Exit(1)
	}

	// Count lines after cleaning
	after, err := countLines(filename)
	if err != nil {
		fmt.Printf("Error counting lines after cleaning: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("The file %s is now %d lines\n", filename, after)
}

func main() {
	// Prompt the user for a file to clean
	var filename string
	fmt.Print("Enter a file to clean: ")
	fmt.Scanln(&filename)

	// Perform the file cleaning
	cleanFile(filename)

	// Exit with code 1 as in the original script
	os.Exit(1)
}

