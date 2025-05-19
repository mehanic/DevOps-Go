package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to print a string in green
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

	// Print the header
	green("     Name:    UID:      Shell:")

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by ':' character (the field separator)
		fields := strings.Split(line, ":")

		// Check if the line has enough fields
		if len(fields) >= 7 {
			// Print the desired fields: Name, UID, and Shell
			name := fields[0] // Username
			uid := fields[2]  // User ID
			shell := fields[6] // User's shell
			fmt.Printf("%10s %4s %17s\n", name, uid, shell)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

