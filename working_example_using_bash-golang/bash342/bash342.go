package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Prompt for file name input
	fmt.Print("Enter the name of the file for reading: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and print each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

