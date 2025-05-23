package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Check for search term argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run search.go <search_term>")
		return
	}
	searchTerm := os.Args[1]

	// Open the input file (you can specify your file here)
	file, err := os.Open("input.txt") // Change "input.txt" to your file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Initialize a buffer to accumulate text
	var record string

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// When we encounter a blank line, we have a record
			if record != "" {
				// Check if the accumulated record matches the search term
				if match, _ := regexp.MatchString(searchTerm, record); match {
					fmt.Println(record) // Print the record if it matches
				}
				// Reset the record for the next block
				record = ""
			}
		} else {
			// Accumulate the line into the record
			if record != "" {
				record += "\n" // Add newline if not the first line of the block
			}
			record += line
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Check if there's a final record to evaluate after the loop
	if record != "" {
		if match, _ := regexp.MatchString(searchTerm, record); match {
			fmt.Println(record) // Print the last record if it matches
		}
	}
}

