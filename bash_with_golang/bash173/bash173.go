package main

import (
	"bufio"
	"fmt"
//	"os"
	"os/exec"
//	"strings"
)

func main() {
	// Command to execute: grep xfs /etc/fstab
	cmd := exec.Command("grep", "xfs", "/etc/fstab")

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
		// For each line, prepend "xfs: "
		line := scanner.Text()
		fmt.Printf("xfs: %s\n", line)
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

