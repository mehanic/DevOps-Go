package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Ensure that at least one username is provided
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <username1> <username2> ...", os.Args[0])
	}

	// Print the script name (similar to $0 in bash)
	fmt.Printf("Executing script: %s\n", os.Args[0])

	// Iterate over the provided users (os.Args[1:] skips the script name)
	for _, user := range os.Args[1:] {
		fmt.Printf("Archiving user: %s\n", user)

		// Lock the account using the `passwd -l` command
		err := runCommand("passwd", "-l", user)
		if err != nil {
			log.Printf("Failed to lock account for user %s: %v", user, err)
			continue // Skip to the next user if the command fails
		}

		// Create a tar archive of the user's home directory
		archivePath := fmt.Sprintf("/archives/%s.tar.gz", user)
		homeDir := fmt.Sprintf("/home/%s", user)
		err = runCommand("tar", "cf", archivePath, homeDir)
		if err != nil {
			log.Printf("Failed to create archive for user %s: %v", user, err)
			continue // Skip to the next user if the command fails
		}

		fmt.Printf("User %s's home directory archived at %s\n", user, archivePath)
	}
}

// runCommand runs a system command and returns an error if it fails.
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

