package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("Received signal and cleaning up files")
	fmt.Println("Cleanup code")
	os.Exit(1)
}

func main() {
	// Create a channel to listen for incoming signals
	sigs := make(chan os.Signal, 1)
	// Notify the channel when the program receives these signals
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGILL)

	// Run the cleanup function when a signal is received
	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %v\n", sig)
		cleanup()
	}()

	// Infinite loop simulating work
	for {
		fmt.Println("Hi there")
		time.Sleep(1 * time.Second)
	}
}

