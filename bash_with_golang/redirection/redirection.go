package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Open the file
	file, err := os.Open("my.file")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize counters
	var lines, words, chars int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		chars += len(line)

		// Count words in the line
		inWord := false
		for _, char := range line {
			if unicode.IsSpace(char) {
				inWord = false
			} else if !inWord {
				inWord = true
				words++
			}
		}
	}

	// Check for errors in the scanner
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print the results similar to `wc`
	fmt.Printf("%d %d %d\n", lines, words, chars)
}

