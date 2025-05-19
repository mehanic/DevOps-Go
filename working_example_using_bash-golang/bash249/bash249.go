package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the filename is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	// Open the input file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line from the file
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by comma
		fields := strings.Split(line, ",")
		if len(fields) < 3 {
			fmt.Println("Invalid line:", line)
			continue
		}

		product := strings.TrimSpace(fields[0])
		price := strings.TrimSpace(fields[1])
		quantity := strings.TrimSpace(fields[2])

		// Print formatted output
		fmt.Printf("\033[1;33m%s ========================\033[0m\n", product)
		fmt.Printf("Price : \t %s \n", price)
		fmt.Printf("Quantity : \t %s \n\n", quantity)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
}

