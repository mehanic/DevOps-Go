package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user to enter a username
	fmt.Print("Enter a user name: ")
	user, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Trim any newlines or spaces from the username
	user = trimNewLine(user)

	// Echo the username
	fmt.Printf("Archiving user: %s\n", user)

	// Lock the account using the `passwd -l` command
	err = runCommand("passwd", "-l", user)
	if err != nil {
		log.Fatalf("Failed to lock the account for user %s: %v", user, err)
	}

	// Create a tar archive of the user's home directory
	archivePath := fmt.Sprintf("/archives/%s.tar.gz", user)
	homeDir := fmt.Sprintf("/home/%s", user)
	err = runCommand("tar", "cf", archivePath, homeDir)
	if err != nil {
		log.Fatalf("Failed to create archive for user %s: %v", user, err)
	}

	fmt.Printf("User %s's home directory archived at %s\n", user, archivePath)
}

// trimNewLine removes newline characters from the user input
func trimNewLine(s string) string {
	return s[:len(s)-1]
}

// runCommand runs a system command and returns an error if it fails
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

