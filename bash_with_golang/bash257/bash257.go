package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if a filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <logfile>")
		return
	}

	logFile := os.Args[1]

	// Open the log file
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a map to hold the count of occurrences for each (9th, 7th) combination
	record := make(map[string]int)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into fields based on whitespace
		fields := strings.Fields(line)

		// Check if there are at least 9 fields (to access the 9th and 7th fields)
		if len(fields) >= 9 {
			key := fmt.Sprintf("%s,%s", fields[8], fields[6]) // Create a key using the 9th and 7th fields
			record[key]++ // Increment the count for this combination
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the occurrences
	for r, count := range record {
		fmt.Printf("%s has occurred %d times.\n", r, count)
	}
}

