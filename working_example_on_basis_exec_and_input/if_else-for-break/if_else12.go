package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Validate age input
	for {
		fmt.Println("Enter your age:")
		ageInput, _ := reader.ReadString('\n')
		ageInput = strings.TrimSpace(ageInput)

		// Check if the input is a valid decimal number
		if _, err := fmt.Sscanf(ageInput, "%d", new(int)); err == nil {
			break
		} else {
			fmt.Println("Please enter a number for your age.")
		}
	}

	// Validate password input (letters and numbers only)
	for {
		fmt.Println("Select a new password (letters and numbers only):")
		passwordInput, _ := reader.ReadString('\n')
		passwordInput = strings.TrimSpace(passwordInput)

		// Regular expression to check if the password is alphanumeric
		if matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, passwordInput); matched {
			break
		} else {
			fmt.Println("Passwords can only have letters and numbers.")
		}
	}
}

// Enter your age:
// 30
// Select a new password (letters and numbers only):
// 34rt
