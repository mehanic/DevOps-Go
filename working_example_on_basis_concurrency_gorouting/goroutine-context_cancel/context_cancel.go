package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done():
		return
	case <-time.After(waitTime):
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}

func main() {
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result)
	}

	foundBy := <-result
	fmt.Println("result found by", foundBy)
	finish()

	time.Sleep(time.Second)
}

// ╰─λ go run context_cancel.go                                                                                                                                         0 (0.001s) < 19:36:48
// 10 sleep 98ms
// 4 sleep 23ms
// 8 sleep 72ms
// 1 sleep 53ms
// 0 sleep 104ms
// 2 sleep 49ms
// 5 sleep 17ms
// 7 sleep 53ms
// 3 sleep 30ms
// 6 sleep 68ms
// 9 sleep 51ms
// worker 5 done
// result found by 5
