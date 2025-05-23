package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var fileSuffix string
	var dirName string

	// Prompt for file suffix
	fmt.Print("Which file types do you want to backup: ")
	fmt.Scanln(&fileSuffix)

	// Prompt for backup directory
	fmt.Print("Which directory do you want to backup to: ")
	fmt.Scanln(&dirName)

	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Create the backup directory if it does not exist
	backupDir := filepath.Join(homeDir, dirName)
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		err = os.MkdirAll(backupDir, 0700) // Create with permission 700
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Walk the home directory
	err = filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Exclude the backup directory from the search
		if path == backupDir {
			return filepath.SkipDir
		}

		// Check if the file matches the specified suffix
		if !info.IsDir() && filepath.Ext(info.Name()) == fileSuffix {
			// Copy the file to the backup directory
			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			destinationFile, err := os.Create(filepath.Join(backupDir, info.Name()))
			if err != nil {
				return err
			}
			defer destinationFile.Close()

			// Copy the file content
			_, err = io.Copy(destinationFile, sourceFile)
			if err != nil {
				return err
			}
			fmt.Printf("Copied %s to %s\n", path, backupDir)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
		return
	}

	fmt.Println("Backup completed successfully.")
	os.Exit(0)
}

