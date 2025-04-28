package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}

// Channel as first-class citizen
// Worker 0 received a
// Worker 1 received b
// Worker 2 received c
// Worker 3 received d
// Worker 4 received e
// Worker 5 received f
// Worker 6 received g
// Worker 7 received h
// Worker 8 received i
// Worker 8 received I
// Worker 4 received E
// Worker 5 received F
// Worker 6 received G
// Worker 7 received H
// Worker 1 received B
// Worker 9 received j
// Worker 9 received J
// Worker 2 received C
// Worker 0 received A
// Worker 3 received D
// Buffered channel
// Worker 0 received a
// Worker 0 received b
// Worker 0 received c
// Worker 0 received d
// Channel close and range
// Worker 0 received a
// Worker 0 received b
// Worker 0 received c
// Worker 0 received d
