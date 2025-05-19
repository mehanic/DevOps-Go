package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./program filename string_length")
		os.Exit(1)
	}

	filename := os.Args[1]
	strLen, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: string_length should be an integer.")
		os.Exit(1)
	}

	// Calculate half of the string length
	count := strLen / 2

	// Build the base pattern using regular expressions
	basePattern := "^"

	for i := 0; i < count; i++ {
		basePattern += "(.)"
	}

	if strLen%2 != 0 {
		basePattern += "."
	}

	for i := count; i > 0; i-- {
		basePattern += "\\" + strconv.Itoa(i)
	}

	basePattern += "$"

	// Compile the generated regex pattern
	regex, err := regexp.Compile(basePattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		os.Exit(1)
	}

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file line by line and apply the regex pattern
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if regex.MatchString(line) {
			fmt.Println(line)
		}
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

