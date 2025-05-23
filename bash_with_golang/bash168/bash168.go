package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the file path
	filePath := "/etc/shadow"

	// Check if the file exists
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("Shadow passwords are enabled.")
	} else if os.IsNotExist(err) {
		fmt.Println("Shadow passwords are not enabled.")
	}

	// Check if the file is writable
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_RDONLY, 0)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("You do NOT have permissions to edit %s.\n", filePath)
		} else {
			fmt.Printf("Error checking permissions for %s: %v\n", filePath, err)
		}
	} else {
		fmt.Printf("You have permissions to edit %s.\n", filePath)
		file.Close() // Close the file when done
	}
}

