package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the script name using filepath.Base to mimic `basename $0`
	scriptName := filepath.Base(os.Args[0])

	// Print the script name and the full path
	fmt.Printf("My name is %s - I was called as %s\n", scriptName, os.Args[0])

	// Get the number of arguments passed to the script (excluding the script name itself)
	numArgs := len(os.Args) - 1
	fmt.Printf("I was called with %d parameters.\n", numArgs)

	// Initialize a counter to mimic the Bash 'count' variable
	count := 1

	// Loop through the arguments starting from os.Args[1] (skipping the script name itself)
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Parameter number %d is: %s\n", count, os.Args[i])
		count++
	}
}

