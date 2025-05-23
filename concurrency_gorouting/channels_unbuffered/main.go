package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) //unbuffered channel

	c2 := make(chan int, 3) //buffered channel

	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine before sending into channel")
		c <- 10
		fmt.Println("func goroutine after sending into channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data...........")
	d := <-c1
	fmt.Println("main goroutine received data: ", d)
	time.Sleep(time.Second)

}

// ─λ go run main.go                                                                        0 (0.001s) < 13:46:58
// main goroutine sleeps for 2 seconds
// func goroutine before sending into channel
// main goroutine starts receiving data...........
// main goroutine received data:  10
// func goroutine after sending into channel
