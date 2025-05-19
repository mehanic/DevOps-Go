package main

import (
	"fmt"
	"time"
)

func main() {
	// This loop runs 10 times.
	for counter := 1; counter <= 10; counter++ {
		fmt.Printf("Hello! This is loop number %d.\n", counter)
		time.Sleep(1 * time.Second) // Sleep for 1 second.
	}

	// After the for-loop finishes, print a goodbye message.
	fmt.Println("All done, thanks for tuning in!")
}

