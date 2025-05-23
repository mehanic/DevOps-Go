package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Check if the file path is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		return
	}

	// Open the file provided as the first argument
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed at the end

	// Declare a slice to hold the lines
	var array []string

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	count := 0

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line) // Append the line to the slice
		count++
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Print the number of elements
	fmt.Printf("Number of elements: %d\n", count)

	// Print the contents of the array
	fmt.Println(array)
}

