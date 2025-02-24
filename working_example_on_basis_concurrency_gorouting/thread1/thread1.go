package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a struct for the custom thread-like functionality
type MyThread struct {
	name      string
	sleepTime time.Duration
}

// Method to run the MyThread logic (similar to the `run` method in Python)
func (mt *MyThread) run(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s start\n", mt.name)
	time.Sleep(mt.sleepTime)
	fmt.Printf("%s end\n", mt.name)
}

func main() {
	// Create a WaitGroup to wait for all threads (goroutines) to finish
	var wg sync.WaitGroup

	// Create instances of MyThread with different sleep durations
	threads := []*MyThread{
		{"Thread-1", 1 * time.Second},
		{"Thread-2", 2 * time.Second},
		{"Thread-3", 3 * time.Second},
	}

	// Start each "thread" (goroutine)
	for _, t := range threads {
		wg.Add(1)
		go t.run(&wg) // Start the thread's logic in a goroutine
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Thread-3 start
// Thread-1 start
// Thread-2 start
// Thread-1 end
// Thread-2 end
// Thread-3 end
