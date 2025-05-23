package main

import (
	"fmt"
	"sync"
	"time"
)

// printTime simulates the Python print_time function
func printTime(threadName string, delay time.Duration, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	for i := 0; i < iterations; i++ {
		time.Sleep(delay * time.Second)
		secondsElapsed := int(time.Since(start).Seconds())
		fmt.Printf("%d %s\n", secondsElapsed, threadName)
	}
}

func main() {
	var wg sync.WaitGroup

	// Add 3 goroutines to WaitGroup
	wg.Add(3)

	// Start the goroutines with different delays and iterations
	go printTime("Fizz", 3, 33, &wg)
	go printTime("Buzz", 5, 20, &wg)
	go printTime("Counter", 1, 100, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}

