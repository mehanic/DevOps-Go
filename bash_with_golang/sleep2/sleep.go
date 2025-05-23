package main

import (
	"fmt"
	"os/exec"
	"time"
)

func repeat(cmd string, args ...string) {
	for {
		// Execute the command
		command := exec.Command(cmd, args...)
		err := command.Run()

		// If the command succeeds, break out of the loop
		if err == nil {
			fmt.Println("Command succeeded")
			return
		}

		// If the command fails, print the error and wait 30 seconds
		fmt.Printf("Command failed: %v\nRetrying in 30 seconds...\n", err)
		time.Sleep(30 * time.Second)
	}
}

func main() {
	// Example usage:
	// This will try to `curl` a website until it succeeds.
	repeat("curl", "http://example.com")
}
