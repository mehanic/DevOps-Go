package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// We need exactly three arguments.
	if len(os.Args) != 4 {
		fmt.Println("Incorrect usage!")
		fmt.Println("Usage: go run file-create.go <directory_name> <file_name> <file_content>")
		os.Exit(1)
	}

	// Save the arguments into variables.
	directoryName := os.Args[1]
	fileName := os.Args[2]
	fileContent := os.Args[3]

	// Create the absolute path for the file.
	absoluteFilePath := filepath.Join(directoryName, fileName)

	// Check if the directory exists; otherwise, try to create it.
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		if err := os.Mkdir(directoryName, 0755); err != nil {
			fmt.Println("Cannot create directory, exiting program!")
			os.Exit(1)
		}
	}

	// Try to create the file, if it does not exist.
	if _, err := os.Stat(absoluteFilePath); os.IsNotExist(err) {
		file, err := os.Create(absoluteFilePath)
		if err != nil {
			fmt.Println("Cannot create file, exiting program!")
			os.Exit(1)
		}
		defer file.Close() // Ensure the file is closed after writing
	}

	// File has been created, write the content to it.
	if err := os.WriteFile(absoluteFilePath, []byte(fileContent), 0644); err != nil {
		fmt.Println("Error writing to file, exiting program!")
		os.Exit(1)
	}

	fmt.Printf("File %s created successfully with the specified content.\n", absoluteFilePath)
}

