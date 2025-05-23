package main

import "fmt"

func GenInt64() chan int64 {
	ch := make(chan int64)
	go func() {
		for i := int64(0); ; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	ch1, ch2 := GenInt64(), GenInt64()
	for i := 0; i < 20; i++ {
		select {
		case v := <-ch1:
			fmt.Println("ch 1", v)
		case v := <-ch2:
			fmt.Println("ch 2", v)
		}
	}
}

// ╰─λ go run channels.go                                                                                                                                               0 (0.002s) < 22:51:31
// ch 2 0
// ch 2 1
// ch 2 2
// ch 2 3
// ch 2 4
// ch 2 5
// ch 2 6
// ch 2 7
// ch 2 8
// ch 2 9
// ch 2 10
// ch 2 11
// ch 2 12
// ch 2 13
// ch 2 14
// ch 2 15
// ch 2 16
// ch 2 17
// ch 2 18
// ch 2 19
