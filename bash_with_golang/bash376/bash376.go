package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the name of the program (equivalent to $0 in Bash)
	programName := os.Args[0]

	// Print the program name
	fmt.Printf("I was ran as: %s\n", programName)
}

