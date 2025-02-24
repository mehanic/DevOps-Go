package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)

	go func(c chan int) {
		for i := 0; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c)

	fmt.Println("Main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c {
		fmt.Println("Main goroutine received value from channel: ", v)
	}
}

// ╰─λ go run main.go                                                                        0 (0.001s) < 13:34:30
// Main goroutine sleeps for 2 seconds
// func goroutine #0 starts sending data into the channel
// func goroutine #0 after sending data into the channel
// func goroutine #1 starts sending data into the channel
// func goroutine #1 after sending data into the channel
// func goroutine #2 starts sending data into the channel
// Main goroutine received value from channel:  0
// Main goroutine received value from channel:  1
// Main goroutine received value from channel:  2
// func goroutine #2 after sending data into the channel
// func goroutine #3 starts sending data into the channel
// func goroutine #3 after sending data into the channel
// func goroutine #4 starts sending data into the channel
// func goroutine #4 after sending data into the channel
// func goroutine #5 starts sending data into the channel
// func goroutine #5 after sending data into the channel
// Main goroutine received value from channel:  3
// Main goroutine received value from channel:  4
// Main goroutine received value from channel:  5
