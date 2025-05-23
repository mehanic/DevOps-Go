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

	// Infinite loop using the condition like 'until'
	for !exitPlease {
		// Print message
		fmt.Println("Press CTRL+C to stop...")

		// Sleep for 1 second (like `sleep 1` in Bash)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Exiting gracefully.")
}

