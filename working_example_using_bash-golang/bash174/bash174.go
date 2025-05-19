package main

import (
	"bufio"
	"fmt"
//	"os"
	"os/exec"
	"strings"
)

// Mock function to simulate checking if a database was backed up recently
func dbBackedUpRecently(db string) bool {
	// Implement your logic to check if the database was backed up recently
	// For demonstration, returning false to indicate it needs a backup
	return false
}

// Mock function to simulate backing up the database
func backup(db string) {
	// Implement your backup logic here
	fmt.Printf("Backing up database: %s\n", db)
}

func main() {
	// Prepare the command to show databases
	cmd := exec.Command("mysql", "-BNe", "show databases")

	// Get the standard output pipe
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error getting stdout pipe: %v\n", err)
		return
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %v\n", err)
		return
	}

	// Create a scanner to read the output
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		db := strings.TrimSpace(scanner.Text()) // Read database name
		if dbBackedUpRecently(db) { // Check if backed up recently
			continue // Skip backup if true
		}
		backup(db) // Otherwise, perform backup
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from command output: %v\n", err)
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Error waiting for command: %v\n", err)
	}
}

