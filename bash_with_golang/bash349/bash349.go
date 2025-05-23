package main

import (
	"fmt"
)

func main() {
	var number1, number2 int

	// Prompt for the first value
	fmt.Print("Enter first value: ")
	_, err := fmt.Scanln(&number1)
	if err != nil {
		fmt.Println("Invalid input for first value")
		return
	}

	// Prompt for the second value
	fmt.Print("Enter second value: ")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Println("Invalid input for second value")
		return
	}

	// Calculate total using addition
	total := number1 + number2
	fmt.Println(total)

	// Calculate sum using a different method (although it's the same in Go)
	sum := number1 + number2
	fmt.Printf("Sum is %d\n", sum)

	// This is the same as the previous sum, as Go does not support the $[] syntax
	// But we can still show the total again
	fmt.Printf("Sum is %d\n", total)
}

