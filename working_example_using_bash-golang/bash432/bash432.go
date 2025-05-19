package main

import (
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// Timer function that runs for the specified duration
func funcTimer(duration time.Duration, done chan struct{}) {
	time.Sleep(duration)
	close(done) // Notify that the timer is done
}

// Cleanup function to handle termination
func cleanUp() {
	// Perform any cleanup actions here
	// For example, you could log a message or clean up resources
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(1) // Ensure at least one argument is provided
	}

	// Parse the timer duration from the first argument
	timerDuration, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1) // Exit if the argument is not a valid number
	}

	// Create a channel to signal timer completion
	timerDone := make(chan struct{})

	// Start the timer in a goroutine
	go funcTimer(time.Duration(timerDuration)*time.Second, timerDone)

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Execute the command passed as arguments
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command
	if err := cmd.Start(); err != nil {
		cleanUp()
		os.Exit(1)
	}

	// Wait for either the timer or a signal
	select {
	case <-timerDone:
		// Timer completed
		cleanUp()
		cmd.Process.Kill() // Terminate the command
	case sig := <-sigChan:
		// Signal received (SIGINT or SIGTERM)
		cleanUp()
		cmd.Process.Kill() // Terminate the command
		// Optionally handle other cleanup logic based on the signal
		_ = sig
	}

	// Wait for the command to finish (if it hasn't been killed)
	if err := cmd.Wait(); err != nil {
		// Handle the error if the command exits with a non-zero status
		os.Exit(1)
	}

	os.Exit(0) // Exit successfully
}

