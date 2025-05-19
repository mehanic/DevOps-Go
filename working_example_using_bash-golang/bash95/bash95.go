package main

import (
	"fmt"
	"time"
)

func main() {
	// Capture the start time
	btime := time.Now()

	// Sleep for 2 seconds
	time.Sleep(2 * time.Second)

	// Capture the end time
	etime := time.Now()

	// Calculate the time difference in seconds
	diff := etime.Sub(btime).Seconds()

	// Print the time difference
	fmt.Printf("Task time: %.0f seconds\n", diff)
}

