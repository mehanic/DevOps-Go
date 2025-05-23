package main

import (
	"bufio"
	"fmt"
	"os"
//	"strings"
)

// Function to read lines from a file and return them as a slice of strings
func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Function to combine two files, line by line, with a specified delimiter
func pasteFiles(file1, file2, delimiter string) error {
	lines1, err := readLinesFromFile(file1)
	if err != nil {
		return err
	}

	lines2, err := readLinesFromFile(file2)
	if err != nil {
		return err
	}

	// Determine the length of the shorter file
	maxLines := len(lines1)
	if len(lines2) > maxLines {
		maxLines = len(lines2)
	}

	// Print the lines from both files with the specified delimiter
	for i := 0; i < maxLines; i++ {
		var line1, line2 string

		// Ensure we don't go out of bounds
		if i < len(lines1) {
			line1 = lines1[i]
		}

		if i < len(lines2) {
			line2 = lines2[i]
		}

		fmt.Println(line1 + delimiter + line2)
	}

	return nil
}

func main() {
	// Default file paths
	file1 := "1.txt"
	file2 := "2.txt"

	// Part 1: Default delimiter (tab character)
	fmt.Println("Pasting with default delimiter (tab):")
	err := pasteFiles(file1, file2, "\t")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Part 2: Custom delimiter (comma)
	fmt.Println("\nPasting with custom delimiter (,):")
	err = pasteFiles(file1, file2, ",")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

