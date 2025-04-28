package main

import "fmt"

func main() {

	jobs := make(chan int, 3)

	jobs <- 1
	jobs <- 2
	jobs <- 3

	close(jobs)

	for {
		j, more := <-jobs

		if !more {
			return
		}

		fmt.Println(j)
		fmt.Println(more)
	}

}

// ╰─λ go run closing_channels.go                                                            0 (0.001s) < 13:57:50
// 1
// true
// 2
// true
// 3
// true
