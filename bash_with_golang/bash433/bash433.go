package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"golang.org/x/sys/unix" // Import the unix package for Mkfifo
)

const (
	fifoFile   = "/tmp/WORK_QUEUE_FIFO"
	numWorkers = 5
)

func main() {
	// Remove existing FIFO file
	if err := os.RemoveAll(fifoFile); err != nil {
		fmt.Println("Error removing FIFO:", err)
		return
	}

	// Create FIFO
	if err := unix.Mkfifo(fifoFile, 0600); err != nil {
		fmt.Println("Error creating FIFO:", err)
		return
	}

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i)
	}

	// Write to the FIFO
	for i := 0; i < numWorkers; i++ {
		f, err := os.OpenFile(fifoFile, os.O_WRONLY|os.O_NONBLOCK, 0600) // Open FIFO for writing
		if err != nil {
			fmt.Println("Error opening FIFO for writing:", err)
			return
		}

		_, err = f.WriteString(strconv.Itoa(i) + "\n")
		if err != nil {
			fmt.Println("Error writing to FIFO:", err)
			return
		}
		f.Close() // Close the file after writing
	}

	// Wait for a few seconds before cleaning up
	time.Sleep(5 * time.Second)

	// Cleanup: Remove FIFO
	if err := os.Remove(fifoFile); err != nil {
		fmt.Println("Error removing FIFO:", err)
	}
}

func worker(id int) {
	for {
		// Open FIFO for reading
		f, err := os.OpenFile(fifoFile, os.O_RDONLY|os.O_NONBLOCK, 0600) // Open FIFO for reading
		if err != nil {
			fmt.Println("Error opening FIFO for reading:", err)
			return
		}

		// Read from the FIFO
		data, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("Error reading from FIFO:", err)
			return
		}

		// Process the data
		if len(data) > 0 {
			number, err := strconv.Atoi(string(data[:len(data)-1])) // Remove newline
			if err != nil {
				fmt.Println("Error converting data to int:", err)
				return
			}
			fmt.Printf("Worker %d processed number: %d\n", id, number)
		}

		f.Close() // Close the file after reading
		time.Sleep(1 * time.Second) // Sleep to avoid busy waiting
	}
}

