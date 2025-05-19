package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup(signal os.Signal) {
	fmt.Println("Received signals and cleanup files")
	fmt.Println("cleanup code")
	fmt.Printf("Signals received is %v\n", signal)
	os.Exit(1)
}

func main() {
	// Create a channel to listen for OS signals
	signalChannel := make(chan os.Signal, 1)

	// Notify the channel on specific signals
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGHUP, syscall.SIGILL)

	// Start a goroutine to handle signals
	go func() {
		sig := <-signalChannel
		cleanup(sig)
	}()

	// Main loop
	for {
		fmt.Println("Hi there")
		time.Sleep(1 * time.Second)
	}
}

