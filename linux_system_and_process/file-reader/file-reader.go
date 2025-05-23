package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "pi_digits.txt"

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var lines []string

	// Read all lines into a slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print each line with trailing whitespace removed
	for _, line := range lines {
		fmt.Println(strings.TrimSpace(line))
	}
}

