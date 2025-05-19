package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to print text in green
func green(s string) {
	fmt.Printf("\033[1;32m%s\033[0m\n", s)
}

func main() {
	// Open the /etc/passwd file
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Print the header in green
	green("  Name:      UID:  Shell:")

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line by the ':' character
		fields := strings.Split(line, ":")

		// Check if there are at least 7 fields
		if len(fields) >= 7 {
			// Print formatted output
			fmt.Printf("%10s %4s %17s\n", fields[0], fields[2], fields[6])
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

