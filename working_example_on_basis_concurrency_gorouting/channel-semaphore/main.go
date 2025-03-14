package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
		close(done)
	}()

	for val := range c {
		fmt.Println(val)
	}
}


// 0
// 1
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
