package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 6
	goroutinesNum = 5
	quotaLimit    = 2
)

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{} // ratelim.go, берём свободный слот
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))

		if j%2 == 0 {
			<-quotaCh             // ratelim.go, возвращаем слот
			quotaCh <- struct{}{} // ratelim.go, берём слот
		}

		runtime.Gosched() // даём поработать другим горутинам
	}
	<-quotaCh // ratelim.go, возвращаем слот
}

func main() {
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit) // ratelim.go
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

// ╰─λ go run ratelim.go                                                                                                                                                0 (0.002s) < 19:43:22
//          █    th 4 iter 0
//          █    th 4 iter 1 ■
//          █    th 4 iter 2 ■■
//      █        th 2 iter 0
//        █      th 3 iter 0
//          █    th 4 iter 3 ■■■
//          █    th 4 iter 4 ■■■■
//      █        th 2 iter 1 ■
//      █        th 2 iter 2 ■■
//        █      th 3 iter 1 ■
//        █      th 3 iter 2 ■■
//    █          th 1 iter 0
//          █    th 4 iter 5 ■■■■■
//        █      th 3 iter 3 ■■■
//        █      th 3 iter 4 ■■■■
//  █            th 0 iter 0
//    █          th 1 iter 1 ■
//    █          th 1 iter 2 ■■
//        █      th 3 iter 5 ■■■■■
//  █            th 0 iter 1 ■
//  █            th 0 iter 2 ■■
//    █          th 1 iter 3 ■■■
//    █          th 1 iter 4 ■■■■
//  █            th 0 iter 3 ■■■
//  █            th 0 iter 4 ■■■■
//      █        th 2 iter 3 ■■■
//      █        th 2 iter 4 ■■■■
//      █        th 2 iter 5 ■■■■■
//    █          th 1 iter 5 ■■■■■
//  █            th 0 iter 5 ■■■■■
