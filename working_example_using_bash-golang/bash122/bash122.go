package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Step 1: Check command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Step 2: Read the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Step 3: Use regex to match words
	re := regexp.MustCompile(`\b\w+\b`)

	// Step 4: Count occurrences of each word
	wordCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := re.FindAllString(line, -1)
		for _, word := range words {
			// Convert word to lowercase to avoid case-sensitive issues
			word = strings.ToLower(word)
			wordCount[word]++
		}
	}

	// Step 5: Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Step 6: Print results in a formatted way
	fmt.Printf("%-14s %s\n", "Word", "Count")
	// Sort the words alphabetically
	var keys []string
	for key := range wordCount {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%-14s %d\n", key, wordCount[key])
	}
}

