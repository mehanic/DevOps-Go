package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the total number of parameters (excluding the script name)
	totalParams := len(os.Args) - 1
	fmt.Printf("Total number of parameters are = %d\n", totalParams)

	// Print the script name (os.Args[0])
	fmt.Printf("Script name = %s\n", os.Args[0])

	// Print individual parameters if they exist
	if totalParams >= 1 {
		fmt.Printf("First Parameter is %s\n", os.Args[1])
	}
	if totalParams >= 2 {
		fmt.Printf("Second Parameter is %s\n", os.Args[2])
	}
	if totalParams >= 3 {
		fmt.Printf("Third Parameter is %s\n", os.Args[3])
	}
	if totalParams >= 4 {
		fmt.Printf("Fourth Parameter is %s\n", os.Args[4])
	}
	if totalParams >= 5 {
		fmt.Printf("Fifth Parameter is %s\n", os.Args[5])
	}

	// Print all parameters
	fmt.Printf("All parameters are = %v\n", os.Args[1:])
}

