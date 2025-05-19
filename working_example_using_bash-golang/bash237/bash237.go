package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	var fileCompression, dirName string

	// Prompt user for compression choice
	fmt.Print("Choose H, M or L compression: ")
	fmt.Scanln(&fileCompression)

	// Prompt user for backup directory
	fmt.Print("Which directory do you want to backup to: ")
	fmt.Scanln(&dirName)

	// Create the backup directory if it does not exist
	backupDir := filepath.Join(os.Getenv("HOME"), dirName)
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		err := os.MkdirAll(backupDir, 0700) // set directory permissions to 700
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Define tar command options based on compression choice
	var tarOpt string
	switch fileCompression {
	case "L":
		tarOpt = fmt.Sprintf("-cvf %s/b.tar --exclude %s %s", backupDir, backupDir, os.Getenv("HOME"))
	case "M":
		tarOpt = fmt.Sprintf("-czvf %s/b.tar.gz --exclude %s %s", backupDir, backupDir, os.Getenv("HOME"))
	case "H":
		tarOpt = fmt.Sprintf("-cjvf %s/b.tar.bzip2 --exclude %s %s", backupDir, backupDir, os.Getenv("HOME"))
	default:
		fmt.Println("Invalid compression option. Please choose H, M, or L.")
		return
	}

	// Execute the tar command
	cmd := exec.Command("tar", append([]string{tarOpt})...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing tar command:", err)
		return
	}

	fmt.Println("Backup completed successfully.")
}

