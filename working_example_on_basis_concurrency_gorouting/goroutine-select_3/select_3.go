package main

import (
	"fmt"
)

func main() {
	cancelCh := make(chan struct{})
	dataCh := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				return
			case dataCh <- val:
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{}
			break
		}
	}
}

// ╰─λ go run select_3.go                                                                                                                                               0 (0.001s) < 19:45:00
// read 0
// read 1
// read 2
// read 3
// read 4
// send cancel
