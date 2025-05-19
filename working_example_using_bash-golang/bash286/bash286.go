package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var fileCompression string
	var dirName string

	// Prompt the user for compression type
	fmt.Print("Choose H, M or L compression: ")
	fmt.Scanln(&fileCompression)

	// Prompt the user for the backup directory
	fmt.Print("Which directory do you want to backup to: ")
	fmt.Scanln(&dirName)

	// Create the backup directory if it does not exist
	backupDir := os.Getenv("HOME") + "/" + dirName
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		if err := os.MkdirAll(backupDir, 0700); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Determine the tar command based on compression level
	var tarCmd string
	switch strings.ToUpper(fileCompression) {
	case "L":
		tarCmd = fmt.Sprintf("tar -cvf %s/b.tar --exclude %s %s", backupDir, backupDir, os.Getenv("HOME"))
	case "M":
		tarCmd = fmt.Sprintf("tar -czvf %s/b.tar.gz --exclude %s %s", backupDir, backupDir, os.Getenv("HOME"))
	default:
		tarCmd = fmt.Sprintf("tar -cjvf %s/b.tar.bzip2 --exclude %s %s", backupDir, backupDir, os.Getenv("HOME"))
	}

	// Execute the tar command
	cmd := exec.Command("bash", "-c", tarCmd)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing tar command:", err)
		return
	}

	fmt.Println("Backup completed successfully.")
}

