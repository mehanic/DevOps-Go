package main

import (
	"fmt"
	"os"
)

// Function to create a file
func createFile(fname string) error {
	// Create the file using os.Create
	file, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close() // Ensure the file is closed after creation

	// Check if the file was created by listing it
	return checkFileExists(fname)
}

// Function to delete a file
func deleteFile(fname string) error {
	// Remove the file using os.Remove
	err := os.Remove(fname)
	if err != nil {
		return fmt.Errorf("error deleting file: %v", err)
	}

	// Check if the file was deleted by listing it
	return checkFileExists(fname)
}

// Function to check if a file exists
func checkFileExists(fname string) error {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		// If the file does not exist, return nil to indicate success for the deletion case
		fmt.Printf("File does not exist: %s\n", fname)
		return nil
	}

	// If the file exists, print its name
	fmt.Printf("File exists: %s\n", fname)
	return nil
}

func main() {
	fname := "my_test_file.txt" // Name of the file

	// Create the file
	if err := createFile(fname); err != nil {
		fmt.Println(err)
		return
	}

	// Delete the file
	if err := deleteFile(fname); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Finished operations successfully.")
}

