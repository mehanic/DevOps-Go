package main

import (
	"fmt"
	"sync"
	"time"
)

// Crawler struct simulates the Crawler class in Python
type Crawler struct {
	done bool
	mu   sync.Mutex
}

// NewCrawler creates and returns a new Crawler instance
func NewCrawler() *Crawler {
	return &Crawler{done: false}
}

// isDone checks if the Crawler has finished its task
func (c *Crawler) isDone() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.done
}

// run simulates the task and error handling (like raising Exception in Python)
func (c *Crawler) run(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	
	time.Sleep(5 * time.Second) // Simulate some task

	// Simulate setting the task as done
	c.mu.Lock()
	c.done = true
	c.mu.Unlock()

	// Simulate an exception being thrown (panic in Go)
	panic("Something bad happened!")
}

func main() {
	// Using sync.WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Create and start a Crawler goroutine
	crawler := NewCrawler()
	wg.Add(1)
	go crawler.run(&wg)

	for {
		time.Sleep(1 * time.Second)

		// Check if the Crawler task is done
		if crawler.isDone() {
			fmt.Println("Done")
			break
		}

		// Check if the goroutine has panicked
		if !crawler.isDone() {
			// If the goroutine has finished (either successfully or via panic), start a new one
			fmt.Println("Restarting Crawler")
			crawler = NewCrawler()
			wg.Add(1)
			go crawler.run(&wg)
		}
	}

	// Wait for all goroutines to complete
	wg.Wait()
}

