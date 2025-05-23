package main

import (
	"log"
	"sync"
	"time"
)

func delayed(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	log.Printf("%s worker running", name)
}

func main() {
	// Set up logging format similar to Python's logging setup
	log.SetFlags(0)

	// Create a WaitGroup to wait for the timers to finish
	var wg sync.WaitGroup

	// Start timers
	log.Println("starting timers")

	// Timer t1
	wg.Add(1)
	time.AfterFunc(3*time.Second, func() {
		delayed(&wg, "t1")
	})

	// Timer t2
	wg.Add(1)
	t2 := time.AfterFunc(3*time.Second, func() {
		delayed(&wg, "t2")
	})

	// Wait for 2 seconds before canceling t2
	log.Println("waiting before canceling t2")
	time.Sleep(2 * time.Second)

	// Cancel timer t2
	log.Println("canceling t2")
	t2.Stop()

	// Wait for the remaining timers to finish
	wg.Wait()
	log.Println("done")
}


// ╰─λ go run thread3.go                                                                     0 (0.001s) < 14:50:19
// starting timers
// waiting before canceling t2
// canceling t2
// t1 worker running
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0xc000124ea0?)
// 	/home/mehanic/.gvm/gos/go1.23.0/src/runtime/sema.go:71 +0x25
// sync.(*WaitGroup).Wait(0xc00012a0f0?)
// 	/home/mehanic/.gvm/gos/go1.23.0/src/sync/waitgroup.go:118 +0x48
// main.main()
// 	/home/mehanic/structure/concurrency_i_goroutines/thread3/thread3.go:45 +0x267
// exit status 2
