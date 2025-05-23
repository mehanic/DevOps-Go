package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	iterationsNum = 4
	goroutinesNum = 4
)

func main() {
	for i := 0; i < goroutinesNum; i++ {
		go func(th int) {
			for j := 0; j < iterationsNum; j++ {
				fmt.Printf(formatWork(th, j))
				time.Sleep(time.Millisecond)
			}
		}(i)
	}
	fmt.Scanln()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

func imports() {
	fmt.Println(time.Millisecond, runtime.NumCPU())
}

// ─λ go run goroutines.go                                                                                                                                             0 (6.992s) < 19:38:41
//    █        th 1 iter 0
//        █    th 3 iter 0
//  █          th 0 iter 0
//      █      th 2 iter 0
//      █      th 2 iter 1 ■
//  █          th 0 iter 1 ■
//        █    th 3 iter 1 ■
//    █        th 1 iter 1 ■
//      █      th 2 iter 2 ■■
//  █          th 0 iter 2 ■■
//    █        th 1 iter 2 ■■
//        █    th 3 iter 2 ■■
//  █          th 0 iter 3 ■■■
//      █      th 2 iter 3 ■■■
//        █    th 3 iter 3 ■■■
//    █        th 1 iter 3 ■■■
