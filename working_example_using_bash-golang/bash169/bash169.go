package main

import (
	"fmt"
//	"io/fs"
	"os"
)

func main() {
	// Read user input for the file or directory path
	var filePath string
	fmt.Print("Enter the path to a file or a directory: ")
	fmt.Scanln(&filePath)

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

	// List the contents of the directory if it's a directory
	if info.IsDir() {
		entries, err := os.ReadDir(filePath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}
		fmt.Println("Contents of the directory:")
		for _, entry := range entries {
			if entry.Type().IsDir() {
				fmt.Printf("[DIR] %s\n", entry.Name())
			} else {
				fmt.Printf("[FILE] %s\n", entry.Name())
			}
		}
	} else {
		// For files, we can show the file information
		fmt.Printf("File details: %s\n", filePath)
	}
}

