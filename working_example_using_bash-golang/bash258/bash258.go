package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Check if a search term is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <search_term> <input_file>")
		return
	}

	searchTerm := os.Args[1]
	inputFile := os.Args[2]

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the entire file and process it by records
	var record string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Process the record when encountering an empty line (record separator)
			processRecord(record, searchTerm)
			record = "" // Reset the record
		} else {
			record += line + "\n" // Accumulate lines for the current record
		}
	}
	// Process the last record if there's no trailing new line
	if record != "" {
		processRecord(record, searchTerm)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// processRecord processes each record and checks for the search term
func processRecord(record, searchTerm string) {
	// Split the record into fields using the regex pattern for FS
	re := regexp.MustCompile("[><]")
	fields := re.Split(record, -1)

	// Check if the record contains the search term
	if strings.Contains(record, searchTerm) {
		// Print the desired fields if they exist
		if len(fields) > 13 {
			fmt.Printf("%s: %s %s: %s %s: %s\n", 
				fields[4], fields[5], fields[8], fields[9], fields[12], fields[13])
		}
	}
}

