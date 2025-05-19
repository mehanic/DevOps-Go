package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the full path of the executable
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable:", err)
		return
	}

	// Get the base name of the executable
	executableName := filepath.Base(executable)

	// Print the name of the executable and the full command
	fmt.Printf("You are using %s\n", executableName)
	fmt.Printf("You are using %s\n", executable)

	// Print all arguments passed to the program
	if len(os.Args) > 1 {
		fmt.Printf("Hello %s\n", os.Args[1:])
	} else {
		fmt.Println("Hello (no arguments provided)")
	}
}

