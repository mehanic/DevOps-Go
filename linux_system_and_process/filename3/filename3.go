package main

import (
	"bufio"
	"fmt"
	"os"
//	"strings"
)

func countWords(filename string) {
	// Try to open the file
	file, err := os.Open(filename)
	if err != nil {
		// Failing silently if file not found
		return
	}
	defer file.Close()

	// Read file contents
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Count the words
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	fmt.Printf("The file %s has about %d words.\n", filename, wordCount)
}

func main() {
	filenames := []string{"alice.txt", "siddhartha.txt", "moby_dick.txt", "little_women.txt"}
	for _, filename := range filenames {
		countWords(filename)
	}
}

