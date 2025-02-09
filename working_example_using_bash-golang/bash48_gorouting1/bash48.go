package main

import (
	"fmt"
	"time"
)

// printTime simulates the Python print_time function
func printTime(threadName *string, delay time.Duration, iterations int) {
	start := time.Now()

	for i := 0; i < iterations; i++ {
		time.Sleep(delay * time.Second)
		secondsElapsed := int(time.Since(start).Seconds())

		if threadName != nil {
			fmt.Println(*threadName)
		} else {
			fmt.Println(secondsElapsed)
		}
	}
}

func main() {
	// Start goroutines
	go printTime(nil, 1, 100)           // Simulating the None/anonymous case
	go printTime(ptr("Fizz"), 3, 33)    // Pass "Fizz" as threadName
	go printTime(ptr("Buzz"), 5, 20)    // Pass "Buzz" as threadName

	// Keep the main function running to allow goroutines to finish
	select {}
}

// ptr is a helper function to convert a string to a pointer
func ptr(s string) *string {
	return &s
}

