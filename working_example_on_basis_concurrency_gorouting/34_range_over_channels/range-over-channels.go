package main

import "fmt"

func main() {

	channel := make(chan int, 2)

	channel <- 1
	channel <- 2

	close(channel)

	for element := range channel {
		fmt.Println(element)
	}

}

// ─λ go run range-over-channels.go                                                         0 (0.001s) < 13:58:11
// 1
// 2
