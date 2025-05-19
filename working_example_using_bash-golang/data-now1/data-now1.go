package main

import (
	"fmt"
	"time"
)

func main() {
	// Record the start time
	start := time.Now()

	// Your code logic here
	// For example:
	performTask()

	// Record the end time
	end := time.Now()

	// Calculate the difference in minutes
	diff := end.Sub(start).Minutes()

	// Print the runtime in minutes
	fmt.Printf("Runtime: %.2f minutes\n", diff)
}

// Example function to simulate some work
func performTask() {
	// Simulate a task taking some time
	time.Sleep(2 * time.Second) // Replace this with actual work
	fmt.Println("Task completed.")
}

