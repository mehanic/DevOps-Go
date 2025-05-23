package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sys/unix" // Import the unix package
)

const fifoFile = "/tmp/WORK_QUEUE_FIFO"

func main() {
	// Print the worker ID
	if len(os.Args) < 2 {
		fmt.Println("Worker ID not provided.")
		return
	}
	workerID := os.Args[1]
	fmt.Printf("WORKER started: %s\n", workerID)

	// Create FIFO if it doesn't exist
	if _, err := os.Stat(fifoFile); os.IsNotExist(err) {
		if err := unix.Mkfifo(fifoFile, 0600); err != nil {
			fmt.Printf("Error creating FIFO: %v\n", err)
			return
		}
	}

	// Open the FIFO for reading
	f, err := os.OpenFile(fifoFile, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Printf("Error opening FIFO: %v\n", err)
		return
	}
	defer f.Close()

	// Handle termination signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Start reading from the FIFO
	go func() {
		for {
			var buffer string
			_, err := fmt.Fscanln(f, &buffer) // Read from the FIFO
			if err != nil {
				fmt.Printf("Error reading from FIFO: %v\n", err)
				break
			}

			if buffer != "" {
				fmt.Printf("Worker %s received: %s\n", workerID, buffer)
				break // Exit after processing the message
			}
		}
	}()

	// Wait for a termination signal
	<-c
	fmt.Println("Worker exiting...")
}

