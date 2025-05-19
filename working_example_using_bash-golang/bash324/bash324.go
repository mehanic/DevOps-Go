package main

import (
	"fmt"
	"os"
)

func main() {
	// All arguments passed are in os.Args (excluding the script name)
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments passed.")
		return
	}

	// Print all arguments
	fmt.Println("All Arguments Passed are as follow:")
	fmt.Println(args)

	// Simulate "shift by one"
	if len(args) > 1 {
		args = args[1:]
		fmt.Println("Shift by one position:")
		fmt.Printf("Value of Positional Parameter $1 after shift: %s\n", args[0])
	} else {
		fmt.Println("Not enough arguments to shift by one.")
	}

	// Simulate "shift by two"
	if len(args) > 2 {
		args = args[2:]
		fmt.Println("Shift by two positions:")
		if len(args) > 0 {
			fmt.Printf("Value of Positional Parameter $1 after two shifts: %s\n", args[0])
		} else {
			fmt.Println("No arguments left after two shifts.")
		}
	} else {
		fmt.Println("Not enough arguments to shift by two.")
	}
}

