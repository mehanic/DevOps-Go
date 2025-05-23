package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Prompt the user for how many lines of /etc/passwd to show
	fmt.Print("How many lines of /etc/passwd would you like to see? ")

	// Read user input
	var input string
	fmt.Scanln(&input)

	// Convert the input to an integer
	showLines, err := strconv.Atoi(input)
	if err != nil || showLines < 1 {
		log.Fatalf("Invalid input: %v", err)
	}

	// Open the /etc/passwd file for reading
	file, err := os.Open("/etc/passwd")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize the line number counter
	lineNum := 1

	// Loop through each line of the file
	for scanner.Scan() {
		// Print the current line
		fmt.Println(scanner.Text())

		// Increment the line number
		lineNum++

		// Break if we've reached the number of lines the user requested
		if lineNum > showLines {
			break
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}

