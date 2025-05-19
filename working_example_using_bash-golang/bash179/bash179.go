package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/etc/fstab") // Open the fstab file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after we're done

	scanner := bufio.NewScanner(file) // Create a new scanner to read the file
	lineNum := 1

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text() // Get the current line text
		fmt.Printf("%d: %s\n", lineNum, line) // Print the line number and the line
		lineNum++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

