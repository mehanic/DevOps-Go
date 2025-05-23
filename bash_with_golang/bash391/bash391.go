package main

import (
	"fmt"
	"os"
)

func main() {
	// Filename variable
	fname := "my_test_file.txt"

	// Create the file
	err := createFile(fname)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// Delete the file
	err = deleteFile(fname)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}

// createFile creates a new file with the specified name
func createFile(fname string) error {
	// Create the file
	file, err := os.Create(fname)
	if err != nil {
		return err // Return the error
	}
	defer file.Close() // Ensure the file is closed after creating it

	fmt.Println("File created:", fname)
	return nil
}

// deleteFile deletes the file with the specified name
func deleteFile(fname string) error {
	err := os.Remove(fname)
	if err != nil {
		return err // Return the error
	}

	fmt.Println("File deleted:", fname)
	return nil
}

