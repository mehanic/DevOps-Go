package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the access log file
	file, err := os.Open("access.log") // Change the filename if necessary
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
			statusCode := fields[8] // 9th field (index 8) is the status code
			resource := fields[6]    // 7th field (index 6) is the resource/URL

			// Check if the status code is "404"
			if statusCode == "404" {
				// Use a composite key of status code and resource to count occurrences
				key := fmt.Sprintf("%s,%s", statusCode, resource)
				record[key]++
			}
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the occurrences
	for r, count := range record {
		fmt.Printf("%s has occurred %d times.\n", r, count)
	}
}

