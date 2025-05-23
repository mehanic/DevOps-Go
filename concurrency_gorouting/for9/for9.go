package main

import (
	"bufio"
	"fmt"
	"os"
//	"strconv"
	"time"
)

func main() {
	// Display the program's instructions.
	fmt.Println("Press Enter to begin. Afterwards, press ENTER to 'click' the stopwatch. Press Ctrl-C to quit.")

	// Wait for user input to start the stopwatch
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n') // Wait for Enter key

	fmt.Println("Started.")
	startTime := time.Now()  // Get the first lap's start time
	lastTime := startTime
	lapNum := 1

	// Create a channel to receive user input
	inputChannel := make(chan string)

	go func() {
		for {
			_, _ = reader.ReadString('\n')
			inputChannel <- "tick"
		}
	}()

	// Start tracking the lap times.
	for {
		select {
		case <-inputChannel:
			lapTime := time.Since(lastTime).Round(time.Second * 100)
			totalTime := time.Since(startTime).Round(time.Second * 100)
			fmt.Printf("Lap #%d: %s (%s)\n", lapNum, totalTime, lapTime)
			lapNum++
			lastTime = time.Now() // Reset the last lap time
		}
	}
}


