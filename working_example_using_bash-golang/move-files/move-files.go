package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get all files with `.sh` extension in the current directory
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file has the .sh extension
		if strings.HasSuffix(info.Name(), ".sh") {
			newName := path + ".bash" // Append .bash to the file name
			err := os.Rename(path, newName) // Rename the file
			if err != nil {
				return err
			}
			fmt.Printf("Renamed: %s -> %s\n", path, newName)
		}
		return nil
	})

	// Handle errors during the file renaming process
	if err != nil {
		fmt.Println("Error:", err)
	}
}

