package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) //
	}()

	for i := range c {
		fmt.Println(i)
	}
}

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
