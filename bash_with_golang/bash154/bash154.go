package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file /etc/passwd for reading
	file, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize line number
	lineNum := 1

	// Loop through each line in the file
	for scanner.Scan() {
		// Print the line number and the line content
		fmt.Printf("%d: %s\n", lineNum, scanner.Text())
		lineNum++ // Increment the line number
	}

	// Check for any errors encountered while scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}

