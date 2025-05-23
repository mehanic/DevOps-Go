package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define the usage message
	usage := "Usage: search file string operation"

	// Check if the number of arguments is exactly 3
	if len(os.Args) != 4 {
		fmt.Println(usage)
		os.Exit(2)
	}

	// Assign arguments to variables
	fileName := os.Args[1]
	searchString := os.Args[2]
	operation := os.Args[3]

	// Check if the file exists
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File does not exist.")
		os.Exit(3)
	}
	defer file.Close()

	// Variables for operation message and search option
	var mesg string
	var opt string

	// Determine operation based on the third argument
	switch strings.ToLower(operation) {
	case "c":
		mesg = fmt.Sprintf("Counting the matches in %s of %s", fileName, searchString)
		opt = "count"
	case "p":
		mesg = fmt.Sprintf("Print the matches of %s in %s", searchString, fileName)
		opt = "print"
	case "d":
		mesg = fmt.Sprintf("Printing all lines but those matching %s from %s", searchString, fileName)
		opt = "invert"
	default:
		fmt.Printf("Could not evaluate %s %s %s\n", fileName, searchString, operation)
		os.Exit(1)
	}

	// Print the operation message
	fmt.Println(mesg)

	// Perform the selected operation
	scanner := bufio.NewScanner(file)
	matchCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		switch opt {
		case "count":
			if strings.Contains(line, searchString) {
				matchCount++
			}
		case "print":
			if strings.Contains(line, searchString) {
				fmt.Println(line)
			}
		case "invert":
			if !strings.Contains(line, searchString) {
				fmt.Println(line)
			}
		}
	}

	// Handle counting case
	if opt == "count" {
		fmt.Printf("Total matches: %d\n", matchCount)
	}

	// Handle any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

