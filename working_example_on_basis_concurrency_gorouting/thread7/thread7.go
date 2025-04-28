package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Initialize a sync.Pool to simulate thread-local storage
var pool = sync.Pool{
	New: func() interface{} {
		return nil
	},
}

// myMethod logs the value stored in the thread-local data
func myMethod(data interface{}) {
	if data == nil {
		log.Println("data does not have a value yet")
	} else {
		log.Printf("data value: %v", data)
	}
}

// MyProcess struct simulates a thread, implementing its behavior in the run function
type MyProcess struct {
	wg *sync.WaitGroup
}

// Run method for MyProcess, which simulates thread-local storage
func (p *MyProcess) run() {
	defer p.wg.Done()

	// Retrieve thread-local data from the pool
	data := pool.Get()
	myMethod(data)

	// Set the thread-local value for this goroutine
	data = map[string]interface{}{"process_id": rand.Intn(1000)}
	myMethod(data)

	// Put the data back into the pool for reuse
	pool.Put(data)
}

func main() {
	// Set up logging
	log.SetFlags(0)

	rand.Seed(time.Now().UnixNano())

	// WaitGroup to synchronize all goroutines
	var wg sync.WaitGroup

	// Start 4 processes (goroutines) simulating threads
	for i := 0; i < 4; i++ {
		wg.Add(1)
		process := MyProcess{wg: &wg}
		go process.run()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All processes completed.")
}

// ─λ go run thread7.go                                                                     0 (0.001s) < 14:42:16
// data does not have a value yet
// data does not have a value yet
// data does not have a value yet
// data value: map[process_id:286]
// data value: map[process_id:140]
// data value: map[process_id:522]
// data does not have a value yet
// data value: map[process_id:471]
// All processes completed.
