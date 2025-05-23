package main

import (
	"fmt"
)

func main() {
	var number1, number2 float64

	// Get the first number from the user
	fmt.Print("Enter First value: ")
	_, err := fmt.Scanln(&number1)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Get the second number from the user
	fmt.Print("Enter Second value: ")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Perform and display the addition
	sum := number1 + number2
	fmt.Printf("Sum: %.2f\n", sum)

	// Perform and display the division, checking for division by zero
	if number2 != 0 {
		division := number1 / number2
		fmt.Printf("Division: %.2f\n", division)
	} else {
		fmt.Println("Error: Cannot divide by zero.")
	}
}

