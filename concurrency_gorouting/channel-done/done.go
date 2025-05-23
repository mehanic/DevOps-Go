package main

import (
	"fmt"
	"sync"
)

func doWork(id int,
	w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(
	id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}

// Worker 9 received j
// Worker 5 received f
// Worker 6 received g
// Worker 7 received h
// Worker 8 received i
// Worker 1 received b
// Worker 0 received a
// Worker 0 received A
// Worker 1 received B
// Worker 2 received c
// Worker 2 received C
// Worker 3 received d
// Worker 3 received D
// Worker 4 received e
// Worker 4 received E
// Worker 9 received J
// Worker 5 received F
// Worker 6 received G
// Worker 7 received H
// Worker 8 received I
