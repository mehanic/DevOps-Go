package main

import (
	"fmt"
)

func main() {
	// Declare a variable to store user input
	var b int

	// Prompt the user to enter a number
	fmt.Println("Enter a number:")
	fmt.Scanln(&b)

	// Variable to store the sum of the loop
	jami := 0

	// Loop from 10 down to 0, decrementing by 1
	for i := 10; i >= 0; i-- {
		jami += i
	}

	// Print the result
	fmt.Printf("The sum of numbers from 10 to 0 is: %d\n", jami)
}

// Enter a number:
// 5
// The sum of numbers from 10 to 0 is: 55
