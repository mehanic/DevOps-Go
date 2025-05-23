package main

import (
	"fmt"
	"time"
)

func main() {
	// Define the counter outside of the loop.
	counter := 0

	// This loop runs until the counter is greater than 9.
	for counter <= 9 {
		counter++ // Increment the counter by 1.
		fmt.Printf("Hello! This is loop number %d.\n", counter)
		time.Sleep(1 * time.Second) // Wait for 1 second.
	}

	// After the loop finishes, print a goodbye message.
	fmt.Println("All done, thanks for tuning in!")
}

