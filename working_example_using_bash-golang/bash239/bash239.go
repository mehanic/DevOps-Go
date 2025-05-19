package main

import (
	"fmt"
	"os"
)

func main() {
	// Output the name of the executable
	fmt.Printf("You are using %s\n", os.Args[0])

	// Check if an argument is provided
	if len(os.Args) > 1 {
		// Print the greeting message
		fmt.Printf("Hello %s\n", os.Args[1])
	}

	// Exit the program
	os.Exit(0)
}

