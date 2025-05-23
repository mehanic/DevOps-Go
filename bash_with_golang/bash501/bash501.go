package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Check for the current number of arguments.
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Usage: go run find.go <file-name>")
		os.Exit(1)
	}

	// Name of the file to search for.
	fileName := os.Args[1]

	// Use the "find" command to search for the file and suppress errors.
	cmd := exec.Command("find", "/", "-name", fileName)
	cmd.Stderr = os.Stderr // Redirect stderr to os.Stderr to suppress error output in the terminal

	// Get the output of the command.
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// Ignore command errors, such as permission denied.
			_ = exitError
		} else {
			fmt.Println("Error executing command:", err)
			os.Exit(1)
		}
	}

	// Print the results.
	results := strings.TrimSpace(string(output))
	if results == "" {
		fmt.Println("File not found.")
	} else {
		fmt.Println(results)
	}
}

