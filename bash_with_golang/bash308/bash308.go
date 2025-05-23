package main

import (
	"fmt"
	"os"
//	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	dir := "."

	// Walk through the directory and process each file
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the filename contains spaces
		if strings.Contains(path, " ") {
			newPath := strings.ReplaceAll(path, " ", "_")
			
			// Rename the file by replacing spaces with underscores
			err := os.Rename(path, newPath)
			if err != nil {
				fmt.Printf("Error renaming file %s: %v\n", path, err)
				return err
			}
			fmt.Printf("Renamed: %s -> %s\n", path, newPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		os.Exit(1)
	}
}

