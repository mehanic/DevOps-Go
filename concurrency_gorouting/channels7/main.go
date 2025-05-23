package main

import (
	"fmt"
)

func main() {
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// unbuffered channel
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10

	// buffered channel
	ch2 := make(chan int, 1)
	ch2 <- 2
	ch2 <- 3
	fmt.Println(<-ch2)
}

// ─λ go run main.go                                                                         0 (0.001s) < 16:06:36
// 10
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
// 	/home/mehanic/structure/concurrency_i_goroutines/go-concurrency/udemy/02-channels/main.go:29 +0xaa
// exit status 2
