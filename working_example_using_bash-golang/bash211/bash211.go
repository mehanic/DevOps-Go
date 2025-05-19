package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Display the program's instructions.
	fmt.Println("Press enter to begin. Afterwards, press ENTER to 'click' the stopwatch. Press Ctrl-C to quit.")
	
	// Wait for user to press enter to begin
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	fmt.Println("Started.")

	// Start timing
	startTime := time.Now()
	lastTime := startTime
	lapNum := 1

	// Channel to capture Ctrl-C interrupt
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	// Start tracking lap times
	for {
		// Non-blocking check for Ctrl-C
		select {
		case <-sigs:
			fmt.Println("\nDone.")
			return
		default:
			// Wait for user to press enter for next lap
			_, _ = reader.ReadString('\n')

			// Calculate lap time and total time
			lapTime := time.Since(lastTime).Seconds()
			totalTime := time.Since(startTime).Seconds()

			// Print lap information
			fmt.Printf("Lap #%d: %.2f (%.2f)\n", lapNum, totalTime, lapTime)

			// Update for next lap
			lapNum++
			lastTime = time.Now()
		}
	}
}

