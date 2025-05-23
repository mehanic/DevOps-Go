package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Function to handle CTRL+C (SIGINT)
func handleInterrupt() {
	// Create a channel to receive OS signals
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT) // Listen for SIGINT (CTRL+C)

	// Goroutine to handle the signal
	go func() {
		<-sigChannel // Wait for the signal
		fmt.Println("\nCTRL+C Detected!... executing Go trap!")
		os.Exit(1) // Exit the program after handling the signal
	}()
}

func main() {
	// Call the signal handler function
	handleInterrupt()

	// Clear the screen (similar to Bash's clear command)
	fmt.Print("\033[H\033[2J")

	// Simulate a loop similar to the Bash for loop
	for a := 1; a <= 10; a++ {
		fmt.Printf("%d/10 to Exit.\n", a)
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}

	fmt.Println("Exit Go Trap Example!!!")
}

