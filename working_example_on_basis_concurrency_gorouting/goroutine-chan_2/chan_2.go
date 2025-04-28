package main

import "fmt"

func main() {
	in := make(chan int)

	go func(out chan int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("Before", i)
			out <- i
			fmt.Println("After", i)
		}
		close(out)
		fmt.Println("Generator finish")
	}(in)

	for i := range in {
		fmt.Println("\tget", i)
	}

	// fmt.Scanln()
}

// ╰─λ go run chan_2.go                                                                                                                                                 0 (0.001s) < 19:28:19
// Before 0
// After 0
// Before 1
// 	get 0
// 	get 1
// After 1
// Before 2
// After 2
// Before 3
// 	get 2
// 	get 3
// After 3
// Before 4
// After 4
// Generator finish
// 	get 4
