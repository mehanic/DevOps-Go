package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	select {
	case val := <-ch1:
		fmt.Println("ch1 val", val)
	case ch2 <- 1:
		fmt.Println("put val to ch2")
	default:
		fmt.Println("default case")
	}
}

// ╰─λ go run select_1.go                                                                                                                                               0 (0.001s) < 19:43:59
// default case

// ╰─λ go run select_1.go                                                                     0 (0.001s) < 15:53:23
// default case
