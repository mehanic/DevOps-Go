package main

import (
	"fmt"
	"strconv"
)

// Function to add two numbers
func getAddition(a, b int) int {
	return a + b
}

func main() {
	// Prompt user for input
	fmt.Print("Enter your first number: ")
	var input1 string
	fmt.Scanln(&input1)
	
	fmt.Print("Enter your second number: ")
	var input2 string
	fmt.Scanln(&input2)

	// Convert inputs to integers
	a, err1 := strconv.Atoi(input1)
	b, err2 := strconv.Atoi(input2)
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Please enter valid numbers.")
		return
	}

	// Call function and print result
	result := getAddition(a, b)
	fmt.Printf("The addition of %d and %d is: %d\n", a, b, result)
}

// Enter your first number: 12
// Enter your second number: 34
// The addition of 12 and 34 is: 46

