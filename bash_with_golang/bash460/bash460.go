package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if there are at least two arguments passed.
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./arguments-parameters <argument1> <argument2>")
		return
	}

	// Get the arguments (os.Args[0] is the program name, so we start from os.Args[1])
	parameter1 := os.Args[1]
	parameter2 := os.Args[2]

	// Print the passed arguments
	fmt.Printf("This is the first parameter, passed as an argument: %s\n", parameter1)
	fmt.Printf("This is the second parameter, also passed as an argument: %s\n", parameter2)
}

