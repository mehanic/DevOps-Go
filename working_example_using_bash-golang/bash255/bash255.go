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

	// Create a map to hold the count of occurrences for each IP address
	ipCount := make(map[string]int)

	// Print header
	fmt.Println("Log access")

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into fields based on whitespace
		fields := strings.Fields(line)

		// Check if there is at least one field (the IP address)
		if len(fields) > 0 {
			ip := fields[0] // First field is the IP address
			ipCount[ip]++   // Increment the count for this IP address
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the occurrences
	for ip, count := range ipCount {
		fmt.Printf("%s has accessed %d times.\n", ip, count)
	}
}

