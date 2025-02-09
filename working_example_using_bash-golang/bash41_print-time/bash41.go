package main

import (
	"fmt"
	"sync"
	"time"
)

// printTime prints the elapsed time based on the specified delay and number of iterations.
func printTime(threadName string, delay time.Duration, iterations int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when done

	start := time.Now() // Get the current time
	for i := 0; i < iterations; i++ {
		time.Sleep(delay) // Wait for the specified delay
		secondsElapsed := int(time.Since(start).Seconds()) // Calculate elapsed seconds
		if threadName != "" {
			fmt.Println(threadName) // Print thread name if available
		} else {
			fmt.Println(secondsElapsed) // Otherwise, print elapsed seconds
		}
	}
}

func main() {
	var wg sync.WaitGroup // Create a wait group to manage goroutines

	// Add goroutines to the wait group
	wg.Add(3) // We are going to start 3 goroutines

	// Start the goroutines
	go printTime("", 1*time.Second, 100, &wg)     // Equivalent to print_time(None, 1, 100)
	go printTime("Fizz", 3*time.Second, 33, &wg) // Equivalent to print_time("Fizz", 3, 33)
	go printTime("Buzz", 5*time.Second, 20, &wg) // Equivalent to print_time("Buzz", 5, 20)

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("Program complete")
}

