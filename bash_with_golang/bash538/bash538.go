package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Clear the screen (this works on Unix-like systems)
	clearScreen()

	// Create a channel to listen for interrupt signals (CTRL+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	// Start a goroutine to handle the signal
	go func() {
		<-sigChan // Wait for a signal
		fmt.Println("CTRL+C Detected! ...executing Go trap!")
		os.Exit(0) // Exit the program
	}()

	// Loop for 10 seconds
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d/10 to Exit.\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Exit Go Trap Example!!!")
}

// clearScreen clears the console screen
func clearScreen() {
	// ANSI escape code to clear the screen
	fmt.Print("\033[H\033[2J")
}

