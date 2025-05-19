package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Execute the "who" command
	cmd := exec.Command("who")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Count the number of lines in the output
	lineCount := countLines(string(output))
	fmt.Printf("Number of users currently logged in: %d\n", lineCount)
}

// Function to count lines in a string
func countLines(output string) int {
	scanner := bufio.NewScanner(strings.NewReader(output))
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}

// Number of users currently logged in: 2
