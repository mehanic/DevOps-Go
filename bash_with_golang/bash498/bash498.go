package main

import (
	"fmt"
	"time"
)

func main() {
	// Define the counter outside of the loop.
	counter := 0

	// This loop runs 10 times.
	for counter < 10 {
		counter++ // Increment the counter by 1.
		fmt.Printf("Hello! This is loop number %d.\n", counter)
		time.Sleep(1 * time.Second) // Wait for 1 second.
	}

	// After the while-loop finishes, print a goodbye message.
	fmt.Println("All done, thanks for tuning in!")
}

