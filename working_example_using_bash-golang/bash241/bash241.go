package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if at least one argument is provided
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <name>\n", os.Args[0])
		os.Exit(1)
	}

	// Print greeting with the first argument
	fmt.Printf("Hello %s\n", os.Args[1])

	// Exit the program successfully
	os.Exit(0)
}

