package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Check the number of arguments received.
	if len(os.Args) != 2 {
		fmt.Println("Incorrect usage, wrong number of arguments.")
		fmt.Println("Usage: go run square-number.go <number>")
		os.Exit(1)
	}

	inputNumber := os.Args[1]

	// Check to see if the input is a number.
	if !isNumeric(inputNumber) {
		fmt.Println("Incorrect usage, wrong type of argument.")
		fmt.Println("Usage: go run square-number.go <number>")
		os.Exit(1)
	}

	// Convert input string to integer
	number, err := strconv.Atoi(inputNumber)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		os.Exit(1)
	}

	// Multiply the input number by itself and return the result.
	squared := number * number
	fmt.Println(squared)
}

// isNumeric checks if the string represents a numeric value.
func isNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

