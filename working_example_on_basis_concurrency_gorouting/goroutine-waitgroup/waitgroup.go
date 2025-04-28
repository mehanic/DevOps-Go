package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 7
	goroutinesNum = 5
)

func startWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done() // wait_2.go уменьшаем счетчик на 1
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func main() {
	wg := &sync.WaitGroup{} // wait_2.go инициализируем группу
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) // wait_2.go добавляем воркер
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)

	// fmt.Scanln()
	wg.Wait() // wait_2.go ожидаем, пока waiter.Done() не приведёт счетчик к 0
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

// ─λ go run waitgroup.go                                                                                                                                              0 (0.001s) < 19:47:00
//          █    th 4 iter 0
//      █        th 2 iter 0
//        █      th 3 iter 0
//          █    th 4 iter 1 ■
//      █        th 2 iter 1 ■
//        █      th 3 iter 1 ■
//          █    th 4 iter 2 ■■
//        █      th 3 iter 2 ■■
//          █    th 4 iter 3 ■■■
//        █      th 3 iter 3 ■■■
//          █    th 4 iter 4 ■■■■
//        █      th 3 iter 4 ■■■■
//    █          th 1 iter 0
//          █    th 4 iter 5 ■■■■■
//  █            th 0 iter 0
//    █          th 1 iter 1 ■
//          █    th 4 iter 6 ■■■■■■
//  █            th 0 iter 1 ■
//        █      th 3 iter 5 ■■■■■
//    █          th 1 iter 2 ■■
//  █            th 0 iter 2 ■■
//        █      th 3 iter 6 ■■■■■■
//  █            th 0 iter 3 ■■■
//    █          th 1 iter 3 ■■■
//    █          th 1 iter 4 ■■■■
//    █          th 1 iter 5 ■■■■■
//    █          th 1 iter 6 ■■■■■■
//      █        th 2 iter 2 ■■
//      █        th 2 iter 3 ■■■
//      █        th 2 iter 4 ■■■■
//      █        th 2 iter 5 ■■■■■
//      █        th 2 iter 6 ■■■■■■
//  █            th 0 iter 4 ■■■■
//  █            th 0 iter 5 ■■■■■
//  █            th 0 iter 6 ■■■■■■
