package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Replace "kjv.txt" with your KJV file path
	kjvFile := "kjv.txt"

	// Open the KJV file
	file, err := os.Open(kjvFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by colon
		parts := strings.SplitN(line, ":", 4)
		if len(parts) < 4 {
			continue // Skip lines that do not have enough parts
		}

		book := parts[0]
		chapter := parts[1]
		verse := parts[2]
		text := parts[3]

		// Get the first word of the text
		firstWord := strings.Split(text, " ")[0]

		// Print the formatted output
		fmt.Printf("%s %s:%s %s\n", book, chapter, verse, firstWord)
	}

	// Check for errors during file reading
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
}

