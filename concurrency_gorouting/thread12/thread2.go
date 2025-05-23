package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a struct similar to the Python class
type MyThread struct {
	name      string
	sleepTime time.Duration
}

// Method to simulate the "run" method in the Python code
func (mt *MyThread) run(wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	fmt.Printf("%s start\n", mt.name)
	time.Sleep(mt.sleepTime)
	fmt.Printf("%s end\n", mt.name)
}

func main() {
	// Create a WaitGroup to synchronize the goroutines
	var wg sync.WaitGroup

	// Create an array of MyThread structs with different sleep times
	threads := []*MyThread{
		{"Thread-1", 1 * time.Second},
		{"Thread-2", 2 * time.Second},
		{"Thread-3", 3 * time.Second},
	}

	// Start each thread (goroutine)
	for _, t := range threads {
		wg.Add(1)          // Increment the WaitGroup counter for each thread
		go t.run(&wg)      // Start the thread's logic in a goroutine
	}

	// Wait for all threads (goroutines) to finish
	wg.Wait()
}

// ─λ go run thread2.go                                                                     0 (0.001s) < 14:50:54
// Thread-3 start
// Thread-2 start
// Thread-1 start
// Thread-1 end
// Thread-2 end
// Thread-3 end
