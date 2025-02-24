package main

import "fmt"

func main() {
	inc := incrementer(10)
	for i := range puller(inc) {
		fmt.Println(i)
	}
}

func incrementer(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}

// 45
