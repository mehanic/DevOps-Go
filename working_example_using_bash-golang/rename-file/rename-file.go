package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the current directory
	dir := "." // You can specify another directory if needed

	// Walk through the files in the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file ends with .bad
		if filepath.Ext(path) == ".bad" {
			// Construct the new file name with .bash extension
			newPath := path[:len(path)-len(".bad")] + ".bash"
			fmt.Printf("Renaming %s to %s\n", path, newPath)

			// Rename the file
			err := os.Rename(path, newPath)
			if err != nil {
				return fmt.Errorf("error renaming %s to %s: %v", path, newPath, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through the directory: %v\n", err)
	}
}

