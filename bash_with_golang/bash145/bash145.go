package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// backupFile creates a backup of the specified file
func backupFile(filePath string) error {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Return an error if the file does not exist
		return fmt.Errorf("file does not exist: %s", filePath)
	}

	// Construct the backup file path
	backupFileName := fmt.Sprintf("%s.%s.%d", 
		filepath.Base(filePath), 
		time.Now().Format("2006-01-02"), 
		os.Getpid())
	backupFilePath := filepath.Join("/tmp", backupFileName)

	fmt.Printf("Backing up %s to %s\n", filePath, backupFilePath)

	// Read the content of the original file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Write the content to the backup file
	if err := ioutil.WriteFile(backupFilePath, content, 0644); err != nil {
		return fmt.Errorf("error writing backup file: %v", err)
	}

	return nil // Return nil if the backup was successful
}

func main() {
	// Call the backupFile function with the specified file
	if err := backupFile("/etc/hosts"); err != nil {
		fmt.Println("Backup failed:", err)
		os.Exit(1) // Exit with a non-zero status if the backup fails
	}

	// Print success message if the backup succeeds
	fmt.Println("Backup succeeded!")
}

