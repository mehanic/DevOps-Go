package main

import (
	"fmt"
	"os"
//	"path/filepath"
)

func main() {
	// Check if a filename was provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path>")
		os.Exit(3) // Exit with code 3 for missing argument
	}

	// Get the file/directory path from the command-line argument
	filePath := os.Args[1]

	// Get file information
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist.\n", filePath)
		os.Exit(2) // Exit with code 2 for non-existing path
	} else if err != nil {
		fmt.Printf("Error checking the file: %v\n", err)
		os.Exit(2) // Exit with code 2 for other errors
	}

	// Check if it's a regular file or a directory
	if info.Mode().IsRegular() {
		fmt.Printf("%s is a regular file.\n", filePath)
		os.Exit(0) // Exit with code 0 for a regular file
	} else if info.IsDir() {
		fmt.Printf("%s is a directory.\n", filePath)
		os.Exit(1) // Exit with code 1 for a directory
	} else {
		fmt.Printf("%s is something other than a regular file or directory.\n", filePath)
		os.Exit(2) // Exit with code 2 for something else
	}
}

