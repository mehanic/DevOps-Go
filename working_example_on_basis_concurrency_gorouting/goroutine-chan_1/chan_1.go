package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)

	go func(in chan int) {
		val := <-in
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan")
	}(ch1)

	ch1 <- 42
	ch1 <- 100500

	fmt.Println("MAIN: after put to chan")
	fmt.Scanln()
}

// ╰─λ go run chan_1.go                                                                                                                                                 0 (9.353s) < 19:27:39
// GO: get from chan 42
// GO: after read from chan
// MAIN: after put to chan
