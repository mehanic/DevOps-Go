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
		if len(fields) >= 9 && fields[8] == "404" {
			// Create a key based on the 9th and 7th fields
			key := fields[8] + " " + fields[6] // 9th field (status) and 7th field (request)
			// Increment the count for the combination of status and request
			record[key]++
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

