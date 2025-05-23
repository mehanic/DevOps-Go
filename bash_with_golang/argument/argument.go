package main

import (
	"fmt"
	"os"
)

func main() {
	NUM_REQUIRED_ARGS := 2
	numArgs := len(os.Args) - 1

	if numArgs < NUM_REQUIRED_ARGS {
		fmt.Printf("Not enough arguments. Call this program with %s <name> <number>\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]   
	number := os.Args[2]

	fmt.Printf("Name: %s, Number: %s\n", name, number)
}
