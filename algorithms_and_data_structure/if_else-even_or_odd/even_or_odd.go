package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	// Prompt the user for input
	fmt.Print("Enter a number, and I'll tell you if it's even or odd: ")
	fmt.Scanln(&input)

	// Convert input to integer
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Check if the number is even or odd
	if number % 2 == 0 {
		fmt.Printf("\nThe number %d is even.\n", number)
	} else {
		fmt.Printf("\nThe number %d is odd.\n", number)
	}
}

// Enter a number, and I'll tell you if it's even or odd: 4

// The number 4 is even.
