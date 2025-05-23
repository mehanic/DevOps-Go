package main

import (
	"fmt"
	"sync"
	"time"
)

var executedCount int
var mu sync.Mutex

// Function for the goroutine
func printTime(threadName string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done
	for count := 0; count < 5; count++ {
		time.Sleep(delay)
		fmt.Printf("%s: %s\n", threadName, time.Now().Format(time.RFC1123))
	}

	// Use a mutex to safely update the shared executedCount variable
	mu.Lock()
	executedCount++
	mu.Unlock()
}

func main() {
	// Create a WaitGroup to wait for both goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2) // We are adding 2 goroutines to wait for

	// Start two goroutines
	go printTime("Thread-1", 2*time.Second, &wg)
	go printTime("Thread-2", 4*time.Second, &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All threads have completed execution")
}

// Thread-1: Fri, 04 Oct 2024 14:54:38 CEST
// Thread-2: Fri, 04 Oct 2024 14:54:40 CEST
// Thread-1: Fri, 04 Oct 2024 14:54:40 CEST
// Thread-1: Fri, 04 Oct 2024 14:54:42 CEST
// Thread-1: Fri, 04 Oct 2024 14:54:44 CEST
// Thread-2: Fri, 04 Oct 2024 14:54:44 CEST
// Thread-1: Fri, 04 Oct 2024 14:54:46 CEST
// Thread-2: Fri, 04 Oct 2024 14:54:48 CEST
// Thread-2: Fri, 04 Oct 2024 14:54:52 CEST
// Thread-2: Fri, 04 Oct 2024 14:54:56 CEST
// All threads have completed execution
