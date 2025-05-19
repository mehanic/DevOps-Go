package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Ensure that a username is provided
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <username>", os.Args[0])
	}

	username := os.Args[1]

	// Print the script name and the user being archived
	fmt.Printf("Executing script: %s\n", os.Args[0])
	fmt.Printf("Archiving user: %s\n", username)

	// Lock the account using the `passwd -l` command
	err := runCommand("passwd", "-l", username)
	if err != nil {
		log.Fatalf("Failed to lock the account for user %s: %v", username, err)
	}

	// Create an archive of the user's home directory using `tar`
	archivePath := fmt.Sprintf("/archives/%s.tar.gz", username)
	homeDir := fmt.Sprintf("/home/%s", username)
	err = runCommand("tar", "cf", archivePath, homeDir)
	if err != nil {
		log.Fatalf("Failed to create archive for user %s: %v", username, err)
	}

	fmt.Printf("User %s's home directory archived at %s\n", username, archivePath)
}

// runCommand is a helper function to run system commands
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

