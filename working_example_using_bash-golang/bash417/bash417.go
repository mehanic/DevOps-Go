package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var exitPlease = false

// Gracefully handles CTRL+C (SIGINT) to stop the loop
func handleInterrupt() {
	// Create a channel to listen for signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	// Wait for the interrupt signal (CTRL+C)
	<-sigChan
	exitPlease = true
	fmt.Println("\nExit signal received.")
}

func main() {
	// Start the interrupt handler in a separate goroutine
	go handleInterrupt()

	for {
		// Print message
		fmt.Println("Press CTRL+C to stop...")

		// Sleep for 1 second (like `sleep 1` in Bash)
		time.Sleep(1 * time.Second)

		// Check the exit condition and break the loop
		if exitPlease {
			break
		}
	}
	
	fmt.Println("Exiting gracefully.")
}

