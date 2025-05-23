package main

import (
	"fmt"
	"os"
)

func main() {
	// Call the function to check if the right number of arguments are provided.
	if err := checkArguments(2, os.Args[1:]); err != nil {
		fmt.Println("Incorrect usage! Usage: <argument1> <argument2>")
		os.Exit(1)
	}

	// Arguments are correct, print them.
	fmt.Printf("Your arguments are: %s and %s\n", os.Args[1], os.Args[2])
}

// checkArguments validates if the number of provided arguments is as expected.
func checkArguments(expected int, args []string) error {
	if len(args) != expected {
		return fmt.Errorf("incorrect number of arguments")
	}
	return nil
}

