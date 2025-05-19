package main

import (
	"fmt"
	//"io/fs"
	"os"
)

func main() {
	// Check if a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <file-or-directory-path>")
		return
	}

	// Get the file path from command-line arguments
	filePath := os.Args[1]

	// Get file information
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Check if it's a regular file or a directory
	if info.Mode().IsRegular() {
		fmt.Printf("%s is a regular file.\n", filePath)
	} else if info.IsDir() {
		fmt.Printf("%s is a directory.\n", filePath)
	} else {
		fmt.Printf("%s is something other than a regular file or directory.\n", filePath)
	}

	// List the details of the file or directory
	if info.IsDir() {
		entries, err := os.ReadDir(filePath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}
		fmt.Println("Directory contents:")
		for _, entry := range entries {
			if entry.Type().IsDir() {
				fmt.Printf("[DIR] %s\n", entry.Name())
			} else {
				fmt.Printf("[FILE] %s\n", entry.Name())
			}
		}
	} else {
		// For files, we can show the file information
		fmt.Printf("Details of the file: %s\n", filePath)
		fmt.Printf("Size: %d bytes\n", info.Size())
		fmt.Printf("Mode: %s\n", info.Mode())
		fmt.Printf("Modified: %s\n", info.ModTime())
	}
}

