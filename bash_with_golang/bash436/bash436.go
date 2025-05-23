package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
//	"time"
)

const lockFile = "/tmp/mylock"

// setup function to create the lock file
func setup() error {
	// Get the current process ID
	pid := os.Getpid()
	tmpFile := fmt.Sprintf("%s.%d", lockFile, pid)

	// Write the PID to the temporary file
	if err := os.WriteFile(tmpFile, []byte(fmt.Sprintf("%d", pid)), 0600); err != nil {
		return fmt.Errorf("error writing temporary lock file: %v", err)
	}

	// Try to create a hard link for the lock file
	if err := os.Link(tmpFile, lockFile); err != nil {
		if os.IsExist(err) {
			content, _ := os.ReadFile(lockFile) // Read existing lock
			fmt.Printf("Locked by %s\n", content)
			os.Remove(tmpFile) // Remove temp file
			return fmt.Errorf("lock already exists")
		}
		return fmt.Errorf("error creating lock file: %v", err)
	}

	fmt.Println("Created temporary lock")
	// Set up signal handling for cleanup
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		cleanup(tmpFile)
	}()

	return nil
}

// cleanup function to remove the lock files
func cleanup(tmpFile string) {
	os.Remove(tmpFile)
	os.Remove(lockFile)
	fmt.Println("Cleanup done")
	os.Exit(1)
}

func main() {
	// Call the setup function and check for errors
	if err := setup(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Door was left unlocked")

	// Keep the program running to keep the lock active
	select {} // Block forever
}

