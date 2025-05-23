package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a scanner to read from stdin or a specific file
	scanner := bufio.NewScanner(os.Stdin)

	// Print the header
	fmt.Printf("%8s %11s\n", "Username", "Login date")
	fmt.Println("====================")

	cnt := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Skip lines containing "Never logged in", or that are empty
		if strings.Contains(line, "Never logged in") || strings.HasPrefix(line, "Username") || strings.HasPrefix(line, "root") || line == "" {
			continue
		}

		// Split the line into fields
		fields := strings.Fields(line)

		// Count valid records
		cnt++

		// Check the number of fields
		if len(fields) == 8 {
			fmt.Printf("%8s %2s %3s %4s\n", fields[0], fields[4], fields[3], fields[7])
		} else {
			fmt.Printf("%8s %2s %3s %4s\n", fields[0], fields[5], fields[4], fields[8])
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Print the footer and the total count
	fmt.Println("====================")
	fmt.Printf("Total Number of Users Processed: %d\n", cnt)
}

