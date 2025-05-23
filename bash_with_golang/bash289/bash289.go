package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Function to check if the provided argument is a file
func isFile(filename string) error {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || info.IsDir() {
		return fmt.Errorf("%s does not seem to be a file", filename)
	}
	return nil
}

// Function to clean the file by removing commented and blank lines
func cleanFile(filename string) error {
	// Check if it's a valid file
	if err := isFile(filename); err != nil {
		return err
	}

	// Count lines before cleaning
	before, err := countLines(filename)
	if err != nil {
		return err
	}
	fmt.Printf("The file %s starts with %d lines\n", filename, before)

	// Read the file content
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Backup the original file by renaming it with a .bak extension
	err = os.Rename(filename, filename+".bak")
	if err != nil {
		return fmt.Errorf("error creating backup: %v", err)
	}

	// Open the file for writing (cleaned version)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Create a scanner to process the file line by line
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		// Ignore commented lines and blank lines
		if !strings.HasPrefix(trimmed, "#") && len(trimmed) > 0 {
			_, err := file.WriteString(line + "\n")
			if err != nil {
				return fmt.Errorf("error writing to file: %v", err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning file: %v", err)
	}

	// Count lines after cleaning
	after, err := countLines(filename)
	if err != nil {
		return err
	}
	fmt.Printf("The file %s is now %d lines\n", filename, after)

	return nil
}

// Function to count the number of lines in a file
func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

func main() {
	// Prompt user to enter a file
	fmt.Print("Enter a file to clean: ")
	reader := bufio.NewReader(os.Stdin)
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	// Clean the file
	err := cleanFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	os.Exit(0)
}

