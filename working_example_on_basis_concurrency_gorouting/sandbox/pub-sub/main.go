package main

import (
	"fmt"
	"sync"
)

type PubSub[T any] struct {
	subscribers []chan T
	mu          sync.RWMutex
	isClosed    bool
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		mu: sync.RWMutex{},
	}
}

func (ps *PubSub[T]) Subscribe() <-chan T {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.isClosed {
		return nil
	}

	sub := make(chan T)
	ps.subscribers = append(ps.subscribers, sub)

	return sub
}

func (ps *PubSub[T]) Publish(value T) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.isClosed {
		return
	}

	for _, sub := range ps.subscribers {
		sub <- value
	}
}

func (ps *PubSub[T]) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.isClosed {
		return
	}

	for _, sub := range ps.subscribers {
		close(sub)
	}

	ps.isClosed = true
}

func main() {
	ps := NewPubSub[string]()
	wg := sync.WaitGroup{}

	sub1 := ps.Subscribe()
	wg.Add(1)
	go subscriber(&wg, sub1, 1)

	sub2 := ps.Subscribe()
	wg.Add(1)
	go subscriber(&wg, sub2, 2)

	ps.Publish("one")
	ps.Publish("two")
	ps.Publish("three")
	ps.Close()

	wg.Wait()
	fmt.Println("completed.")
}

func subscriber(wg *sync.WaitGroup, sub <-chan string, id int) {
	for {
		val, ok := <-sub
		if !ok {
			fmt.Printf("sub %d, exiting\n", id)
			wg.Done()
			return
		}
		fmt.Printf("sub %d, value %s\n", id, val)
	}
}

// ─λ go run main.go                                                                                                               1 (0.063s) < 14:08:00
// sub 1, value one
// sub 1, value two
// sub 2, value one
// sub 2, value two
// sub 2, value three
// sub 2, exiting
// sub 1, value three
// sub 1, exiting
// completed.
