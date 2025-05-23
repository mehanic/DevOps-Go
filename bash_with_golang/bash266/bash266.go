package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the number of arguments passed to the script (excluding the script name itself)
	numArgs := len(os.Args) - 1

	// Print the number of parameters
	fmt.Printf("I was called with %d parameters.\n", numArgs)

	// Get the script name using filepath.Base to mimic `basename $0`
	scriptName := filepath.Base(os.Args[0])

	// Print the script name
	fmt.Printf("My name is %s - I was called as %s\n", scriptName, os.Args[0])

	// Print up to the first 10 parameters, or show (none) if not available
	for i := 1; i <= 10; i++ {
		if i < len(os.Args) {
			fmt.Printf("My %dth parameter is: %s\n", i, os.Args[i])
		} else {
			fmt.Printf("My %dth parameter is: (none)\n", i)
		}
	}
}

