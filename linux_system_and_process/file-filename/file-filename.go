package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main() {
	filename := "alice.txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Sorry, the file %s does not exist.\n", filename)
		return
	}
	defer file.Close()

	// Read file content and count words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // Split on words
	wordCount := 0

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	} else {
		fmt.Printf("The file %s has about %d words.\n", filename, wordCount)
	}
}
