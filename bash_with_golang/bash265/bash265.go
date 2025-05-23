package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the script name using filepath.Base to mimic `basename $0`
	scriptName := filepath.Base(os.Args[0])

	// Print the script name
	fmt.Printf("My name is %s - I was called as %s\n", scriptName, os.Args[0])

	// Check if at least 2 parameters are passed
	if len(os.Args) > 1 {
		fmt.Printf("My first parameter is: %s\n", os.Args[1])
	} else {
		fmt.Println("My first parameter is: (none)")
	}

	if len(os.Args) > 2 {
		fmt.Printf("My second parameter is: %s\n", os.Args[2])
	} else {
		fmt.Println("My second parameter is: (none)")
	}
}

