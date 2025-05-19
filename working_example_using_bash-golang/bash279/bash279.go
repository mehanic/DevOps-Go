package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Read file suffix and directory name from user
	var fileSuffix, dirName string
	fmt.Print("Which file types do you want to backup: ")
	fmt.Scan(&fileSuffix)
	fmt.Print("Which directory do you want to backup to: ")
	fmt.Scan(&dirName)

	// Create the backup directory if it does not exist
	backupDir := filepath.Join(homeDir, dirName)
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		if err := os.MkdirAll(backupDir, 0700); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Copy matching files
	err = filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the backup directory
		if path == backupDir {
			return nil
		}

		// Check if the file matches the suffix
		if !info.IsDir() && strings.HasSuffix(info.Name(), fileSuffix) {
			destination := filepath.Join(backupDir, info.Name())
			// Copy file to backup directory
			if err := copyFile(path, destination); err != nil {
				fmt.Println("Error copying file:", err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, input, 0644)
}

