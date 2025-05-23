package main

import (
	"fmt"
)

func main() {
	var x int

	// Prompt the user for input
	fmt.Print("Please enter a value: ")
	_, err := fmt.Scan(&x) // Read the user input into x
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Check if the number is even or odd
	if x%2 == 0 {
		fmt.Println("Entered number is even")
	} else {
		fmt.Println("Entered number is odd")
	}
}

