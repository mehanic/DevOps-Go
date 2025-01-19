package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// GetStrings reads value strings from each row of the file
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return lines, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func main() {
	// Path to the file to read
	fileName := "example.txt"

	// Call the GetStrings function to read the file
	lines, err := GetStrings(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Print each line
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Hello, World!
// This is a test file.
// It contains multiple lines of text.
