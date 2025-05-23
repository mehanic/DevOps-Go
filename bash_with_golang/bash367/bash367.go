package main

import (
	"fmt"
	"math"
	"strconv"
	"os"
	"strings"
)

func main() {
	var userInput string
	fmt.Print("Enter input: ")
	_, err := fmt.Scanln(&userInput) // Read user input
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Calculate result with 2 digits after decimal point
	resultWithTwoDigits := roundToDecimal(userInput, 2)
	fmt.Printf("Result with 2 digits after decimal point: %.2f\n", resultWithTwoDigits)

	// Calculate result with 10 digits after decimal point
	resultWithTenDigits := roundToDecimal(userInput, 10)
	fmt.Printf("Result with 10 digits after decimal point: %.10f\n", resultWithTenDigits)

	// Calculate result as rounded integer
	roundedInteger := roundToInteger(userInput)
	fmt.Printf("Result as rounded integer: %d\n", roundedInteger)
}

// roundToDecimal takes a string input and a precision, returns the rounded float64 value
func roundToDecimal(input string, precision int) float64 {
	// Convert input to float64
	value, err := parseInput(input)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return 0
	}
	factor := math.Pow(10, float64(precision))
	return math.Round(value*factor) / factor
}

// roundToInteger takes a string input and returns the rounded integer value
func roundToInteger(input string) int {
	value, err := parseInput(input)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return 0
	}
	return int(math.Round(value))
}

// parseInput tries to convert the input string to float64
func parseInput(input string) (float64, error) {
	// Remove whitespace from input
	input = strings.TrimSpace(input)
	// Parse input as float
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

