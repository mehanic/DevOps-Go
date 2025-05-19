package main

import (
	"fmt"
	"os"
)

func main() {
	var name string

	// Read input from the user
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1) // Set a failed return code
	}

	// Check if the name is empty
	if name == "" {
		fmt.Fprintln(os.Stderr, "No name entered")
		os.Exit(1) // Set a failed return code
	}

	fmt.Println("Name entered:", name)
}

