package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the directory path
	dirPath := "/temporary_dir"

	// Create a new top-level directory
	err := os.Mkdir(dirPath, 0755) // Set appropriate permissions
	if err != nil {
		fmt.Println("mkdir did not successfully complete, stop script execution!")
		os.Exit(1)
	}

	// Create a new file in the temporary directory
	filePath := dirPath + "/tempfile.txt"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close() // Ensure the file is closed when done

	fmt.Println("Directory and file created successfully.")
}

