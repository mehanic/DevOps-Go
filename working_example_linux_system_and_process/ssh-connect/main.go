package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

func sshDeleteLineWithString(host, username, password, filePath, targetString string) error {
	// Set up the SSH client configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // This is insecure for demo purposes
	}

	// Connect to the SSH server
	addr := fmt.Sprintf("%s:%d", host, 22)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}
	defer client.Close()

	// Start a session
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Open the remote file
	rc, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %v", err)
	}
	if err := session.Run(fmt.Sprintf("cat %s", filePath)); err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	// Read the file content
	scanner := bufio.NewScanner(rc)
	var newLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(strings.ToLower(line), strings.ToLower(targetString)) {
			newLines = append(newLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from remote file: %v", err)
	}

	// Write the updated content back to the file
	tempFilePath := filePath + ".tmp"
	err = os.WriteFile(tempFilePath, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("failed to write updated content to temp file: %v", err)
	}

	// Replace the original file with the updated file
	if err := session.Run(fmt.Sprintf("mv %s %s.bak && mv %s %s", filePath, filePath, tempFilePath, filePath)); err != nil {
		return fmt.Errorf("failed to replace the file: %v", err)
	}

	fmt.Printf("Lines containing '%s' (case-insensitive) have been deleted from %s!\n", targetString, filePath)
	return nil
}

func main() {
	// Usage
	host := "your_target_server_ip"
	username := "your_username"
	password := "your_password"
	filePath := "/home/nishanth/Desktop/tobedeleted/delete.txt"
	targetString := "casb"

	err := sshDeleteLineWithString(host, username, password, filePath, targetString)
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
}
