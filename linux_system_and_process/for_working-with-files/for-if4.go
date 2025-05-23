package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "pi_million_digits.txt"

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read lines from the file
	var piString strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		piString.WriteString(strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Get the piString as a string
	piStringValue := piString.String()

	// Print the first 52 characters followed by "..."
	fmt.Println(piStringValue[:52] + "...")
	// Print the length of the piString
	fmt.Println(len(piStringValue))
}

