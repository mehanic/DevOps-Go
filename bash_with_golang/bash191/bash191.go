package main

import (
	"fmt"
	"regexp"
)

func main() {
	var age string
	var password string

	// Loop until a valid age is provided
	for {
		fmt.Println("Enter your age:")
		fmt.Scanln(&age)

		// Check if age is a valid decimal number
		if isDecimal(age) {
			break
		}
		fmt.Println("Please enter a number for your age.")
	}

	// Loop until a valid password is provided
	for {
		fmt.Println("Select a new password (letters and numbers only):")
		fmt.Scanln(&password)

		// Check if password is alphanumeric
		if isAlphanumeric(password) {
			break
		}
		fmt.Println("Passwords can only have letters and numbers.")
	}
}

// Function to check if a string is a decimal number
func isDecimal(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}

// Function to check if a string is alphanumeric
func isAlphanumeric(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(s)
}

