package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("awk.sh")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line) // Split the line into words

		// Process each word
		for _, word := range words {
			// Print each character of the word on a new line
			for i := 0; i < len(word); i++ {
				fmt.Println(string(word[i]))
			}
			fmt.Println() // Print a new line after each word
		}

		// Sleep for 1 second before processing the next line
		time.Sleep(1 * time.Second)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

