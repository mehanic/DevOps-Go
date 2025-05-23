package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string

	// Check if there are command-line arguments
	if len(os.Args) < 2 {
		// No arguments, prompt the user for input
		fmt.Print("Enter a name: ")
		reader := bufio.NewReader(os.Stdin)
		name, _ = reader.ReadString('\n')

		// Remove newline character from input
		name = name[:len(name)-1] // Trim the newline character
	} else {
		// Use the first command-line argument as the name
		name = os.Args[1]
	}

	// Print the greeting message
	fmt.Printf("Hello %s\n", name)
}

