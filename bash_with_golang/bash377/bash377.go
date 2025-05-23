package main

import (
	"fmt"
	"os"
	"path/filepath"
//	"strings"
)

// Global variable to track the current level
var currentLevel int = 0

// Function to create tabs for indentation based on the current level
func tabCreator(level int) string {
	tabs := "."
	for i := 0; i < level; i++ {
		tabs += tabs // doubling the tabs string
	}
	return tabs
}

// Recursive function to print directory tree
func recursiveTree(entry string) {
	// Walk through the directory
	files, err := os.ReadDir(entry)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		leaf := filepath.Join(entry, file.Name())

		// If it's a directory, recursively print its contents
		if file.IsDir() {
			tabs := tabCreator(currentLevel)
			fmt.Printf("%s\\_ %s\n", tabs, leaf)
			currentLevel++
			recursiveTree(leaf)
			currentLevel--
		} else if file.Type().IsRegular() {
			// If it's a file, print with a bar symbol
			tabs := tabCreator(currentLevel)
			fmt.Printf("%s|_ %s\n", tabs, leaf)
		}
	}
}

func main() {
	// Get the directory to walk through from command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory to walk through.")
		return
	}

	parentDir := os.Args[1]

	// Start the recursive function
	recursiveTree(parentDir)
}

