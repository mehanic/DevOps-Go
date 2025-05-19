package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the /etc/passwd file
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println("Error opening /etc/passwd:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop over each line in /etc/passwd
	for scanner.Scan() {
		// Split the line by colon (':') to get fields
		fields := strings.Split(scanner.Text(), ":")
		if len(fields) > 0 {
			// The first field is the username
			username := fields[0]
			fmt.Printf("Username is bss_%s\n", username)
		}
	}

	// Handle any error that occurred during the scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading /etc/passwd:", err)
	}
}

