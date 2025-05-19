package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define the tar command and its arguments
	cmd := exec.Command("tar", "cvf", "home_backup.tar", os.Getenv("HOME"))

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error creating backup: %s\n", err)
		return
	}

	fmt.Println("Backup created successfully: home_backup.tar")
}

