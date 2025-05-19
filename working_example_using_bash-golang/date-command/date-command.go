package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time in seconds
	start := time.Now()

	// Execute the command (in this case, just print "hello world")
	fmt.Println("hello world")

	// Calculate the duration
	end := time.Now()
	difference := end.Sub(start)

	// Output the time taken
	fmt.Printf("Time taken to execute \"echo hello world\" is %.2f seconds.\n", difference.Seconds())
}

