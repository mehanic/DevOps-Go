package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the name of the executable (script)
	scriptName := filepath.Base(os.Args[0])

	// Print the name of the script
	fmt.Printf("You are using %s\n", scriptName)

	// Print the greeting with all arguments
	if len(os.Args) > 1 {
		fmt.Printf("Hello %s\n", os.Args[1:]) // os.Args[1:] to get all arguments
	} else {
		fmt.Println("Hello")
	}

	// Exit with a success status code
	os.Exit(0)
}

