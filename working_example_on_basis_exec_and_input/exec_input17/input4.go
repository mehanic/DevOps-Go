package main

import (
	"fmt"
	"strconv"
)

// Function to multiply a number by 10
func multiplyNum10(value int) int {
	return value * 10
}

func main() {
	// Prompt user for input
	fmt.Print("Enter your number: ")
	var input string
	fmt.Scanln(&input)

	// Convert input to integer
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Call function and print result
	result := multiplyNum10(num)
	fmt.Printf("The result is: %d\n", result)
}

// Enter your number: 12
// The result is: 120
