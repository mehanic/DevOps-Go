package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)
//This function simply calls copyFile to copy a file from sourceFilePath to targetFilePath
func version1(sourceFilePath, targetFilePath string) error {
	// Simply copy the file
	return copyFile(sourceFilePath, targetFilePath)
}

//This function attempts to copy the file using copyFile.
//If the copy operation fails because the target directory does not exist, it creates
// the necessary directory structure using os.MkdirAll.
//After creating the directory, it retries the copy operation.
//If any other error occurs, it prints the error message.

func version2(sourceFilePath, targetFilePath string) {
	err := copyFile(sourceFilePath, targetFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create the parent directory if it doesn't exist
			err := os.MkdirAll(filepath.Dir(targetFilePath), os.ModePerm)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
			// Retry the copy
			err = copyFile(sourceFilePath, targetFilePath)
			if err != nil {
				fmt.Println("Error copying file:", err)
			}
		} else {
			fmt.Println("Error:", err)
		}
	}
}

// copyFile is a helper function to copy files from source to target
func copyFile(sourceFilePath, targetFilePath string) error {
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return err
	}

	// Ensure the target file has the same permissions as the source
	sourceInfo, err := os.Stat(sourceFilePath)
	if err != nil {
		return err
	}
	return os.Chmod(targetFilePath, sourceInfo.Mode())
}

func main() {
	// Define source and target paths (adjust these as needed)
	sourcePath := filepath.Join(os.Getenv("HOME"), "Documents/Writing/Python Cookbook/source")
	targetPath := filepath.Join(os.Getenv("HOME"), "Dropbox/B05442/demo")

	// Find all .rst files recursively in the source path
	err := filepath.Walk(sourcePath, func(sourceFilePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process .rst files
		if filepath.Ext(sourceFilePath) == ".rst" {
			// Determine the relative path and target path
			relativePath, err := filepath.Rel(sourcePath, sourceFilePath)
			if err != nil {
				return err
			}
			targetFilePath := filepath.Join(targetPath, relativePath)

			// Use version2 to copy the file
			version2(sourceFilePath, targetFilePath)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}

