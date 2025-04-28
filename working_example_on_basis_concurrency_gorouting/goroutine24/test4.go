package main

import "fmt"

func main() {
	ch1 := make(chan string, 2)
	ch1 <- "hello"
	ch1 <- "bello"
	close(ch1)
	for val := range ch1 {
		fmt.Println(val)
	}
}

// ╰─λ go run test4.go                                                                       0 (0.001s) < 15:48:22
// hello
// bello
