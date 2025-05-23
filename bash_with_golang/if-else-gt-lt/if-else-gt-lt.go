package main

import (
	"fmt"
	"os"
)

func main() {
	var number int

	// Prompt the user to enter a number
	fmt.Print("Enter a number no greater than 10: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input:", err)
		os.Exit(1) // Set a failed return code
	}

	// Check if the number is greater than 10
	if number > 10 {
		fmt.Fprintf(os.Stderr, "%d is too big\n", number)
		os.Exit(1) // Set a failed return code
	} else {
		fmt.Printf("You entered %d\n", number)
	}
}

