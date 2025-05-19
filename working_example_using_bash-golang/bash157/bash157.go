package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Ensure that a username is passed as an argument
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <username>", os.Args[0])
	}

	username := os.Args[1]

	// Print the script name and user being archived
	fmt.Printf("Executing script: %s\n", os.Args[0])
	fmt.Printf("Archiving user: %s\n", username)

	// Lock the account
	err := runCommand("passwd", "-l", username)
	if err != nil {
		log.Fatalf("Failed to lock account for user %s: %v", username, err)
	}

	// Create a tar archive of the user's home directory
	archivePath := fmt.Sprintf("/archives/%s.tar.gz", username)
	homeDir := fmt.Sprintf("/home/%s", username)

	err = runCommand("tar", "cf", archivePath, homeDir)
	if err != nil {
		log.Fatalf("Failed to create archive for user %s: %v", username, err)
	}

	fmt.Printf("User %s's home directory archived at %s\n", username, archivePath)
}

// runCommand runs a system command with arguments and returns any error.
func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

