package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the directory path
	dirPath := "/temporary_dir"

	// Create a new top-level directory
	if err := os.Mkdir(dirPath, 0755); err != nil {
		fmt.Println("mkdir did not successfully complete, stop script execution!")
		os.Exit(1)
	}

	// Create a new file in our temporary directory
	filePath := dirPath + "/tempfile.txt"
	if _, err := os.Create(filePath); err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Directory and file created successfully.")
}

