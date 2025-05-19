package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the access log file
	file, err := os.Open("access.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a map to hold the count of occurrences
	record := make(map[string]int)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into fields based on whitespace
		fields := strings.Fields(line)

		// Check if there are at least 9 fields
		if len(fields) >= 9 {
			statusCode := fields[8] // Status code is the 9th field (index 8)
			resource := fields[6]    // Resource/URL is the 7th field (index 6)

			// If status code is "404", increment the count for this combination
			if statusCode == "404" {
				record[statusCode+","+resource]++
			}
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

