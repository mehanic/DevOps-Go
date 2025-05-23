package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// choice function that mimics the Bash choice function
func choice(prompt string) string {
	fmt.Print(prompt)

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')

	// Trim whitespace and newline characters
	return strings.TrimSpace(answer)
}

func main() {
	// Simulate the original THISPACKAGE value
	THISPACKAGE := "2024-10-20" // Example date

	CHOICE := ""

	// Until the user confirms with "y"
	for CHOICE != "y" {
		// Display the current package date
		fmt.Printf("This package's date is %s\n", THISPACKAGE)

		// Ask if the date is correct or allow input of a new date
		CHOICE = choice("Is that correct? [Y/,<New date>]: ")

		// If CHOICE is empty, assume "y"
		if CHOICE == "" {
			CHOICE = "y"
		} else if CHOICE != "y" {
			// Override THISPACKAGE with the new date entered by the user
			fmt.Printf("Overriding %s with %s\n", THISPACKAGE, CHOICE)
			THISPACKAGE = CHOICE
		}
	}

	fmt.Println("Confirmed date:", THISPACKAGE)
}

