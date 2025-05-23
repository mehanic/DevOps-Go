package main

import (
	"fmt"
	"os"
)

// Function to create a file
func createFile(fname string) error {
	file, err := os.Create(fname) // Create the file
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed after creating
	fmt.Printf("File created: %s\n", fname)
	return nil
}

// Function to delete a file
func deleteFile(fname string) error {
	err := os.Remove(fname) // Delete the file
	if err != nil {
		return err
	}
	fmt.Printf("File deleted: %s\n", fname)
	return nil
}

func main() {
	fname := "my_test_file.txt" // Name of the file

	// Create the file
	if err := createFile(fname); err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}

	// Delete the file
	if err := deleteFile(fname); err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return
	}

	// Exit successfully
	fmt.Println("Finished operations successfully.")
}

