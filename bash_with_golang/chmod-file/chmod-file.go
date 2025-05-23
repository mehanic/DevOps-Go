package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if any arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: chmod_all file1 file2 ...")
		return
	}

	// Iterate over each argument (file)
	for _, filename := range os.Args[1:] {
		fmt.Printf("Changing permissions on %s\n", filename)

		// Change permissions to 0750
		err := os.Chmod(filename, 0750)
		if err != nil {
			fmt.Printf("Error changing permissions on %s: %v\n", filename, err)
		}
	}
}

