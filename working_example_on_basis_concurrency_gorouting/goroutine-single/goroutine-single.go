package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("this is example with single goroutines ")
	go function()
	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(10 * time.Second)
}

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

// ─λ go run goroutine-single.go                                                            0 (0.001s) < 15:41:41
// this is example with single goroutines
// 10 11 12 13 14 15 16 17 18 19 0123456789
