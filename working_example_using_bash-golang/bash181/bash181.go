package main

import (
	"fmt"
	"os"
)

func main() {
	for index := 1; index < 6; index++ {
		// Creating the project directory name
		projectDir := fmt.Sprintf("/usr/local/project-%d", index)

		// Print the directory being created
		fmt.Printf("Creating %s\n", projectDir)

		// Create the directory
		err := os.Mkdir(projectDir, 0755) // Use 0755 permissions
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", projectDir, err)
		}
	}
}

