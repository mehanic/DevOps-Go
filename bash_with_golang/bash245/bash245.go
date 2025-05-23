package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the base name of the script (equivalent to basename $0)
	scriptName := filepath.Base(os.Args[0])
	fmt.Printf("You are using %s\n", scriptName)

	// Loop over all command-line arguments (skip the first one, which is the script name)
	for _, name := range os.Args[1:] {
		fmt.Printf("Hello %s\n", name)
	}

	// Exit the program successfully
	os.Exit(0)
}

