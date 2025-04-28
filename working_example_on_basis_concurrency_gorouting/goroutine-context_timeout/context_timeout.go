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
	workTime := 50 * time.Millisecond
	ctx, _ := context.WithTimeout(context.Background(), workTime)
	result := make(chan int, 1)

	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result)
	}

	totalFound := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case foundBy := <-result:
			totalFound++
			fmt.Println("result found by", foundBy)
		}
	}
	fmt.Println("totalFound", totalFound)
	time.Sleep(time.Second)
}

// ╰─λ go run context_timeout.go                                                                                                                                        0 (0.013s) < 19:37:29
// 10 sleep 105ms
// 1 sleep 85ms
// 5 sleep 80ms
// 0 sleep 36ms
// 6 sleep 84ms
// 8 sleep 46ms
// 9 sleep 95ms
// 7 sleep 67ms
// 2 sleep 10ms
// 4 sleep 35ms
// 3 sleep 14ms
// worker 2 done
// result found by 2
// worker 3 done
// result found by 3
// worker 4 done
// result found by 4
// worker 0 done
// result found by 0
// worker 8 done
// result found by 8
// totalFound 5
