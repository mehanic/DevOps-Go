package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched() // попробуйте закомментировать
	}
	printFinishWork(workerNum)
}

func main() {
	runtime.GOMAXPROCS(0)               // попробуйте с 0 (все доступные) и 1
	worketInput := make(chan string, 2) // попробуйте увеличить размер канала
	for i := 0; i < goroutinesNum; i++ {
		go startWorker(i, worketInput)
	}

	months := []string{"Январь", "Февраль", "Март",
		"Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь",
		"Октябрь", "Ноябрь", "Декабрь",
	}

	for _, monthName := range months {
		worketInput <- monthName
	}
	close(worketInput) // попробуйте закомментировать

	time.Sleep(time.Millisecond)
}

func formatWork(in int, input string) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"th", in,
		"recieved", input)
}

func printFinishWork(in int) {
	fmt.Println(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", goroutinesNum-in),
		"===", in,
		"finished")
}

// ╰─λ go run workerpool.go                                                                   0 (0.001s) < 15:51:22
//      █    th 2 recieved Январь
//    █      th 1 recieved Март
//  █        th 0 recieved Февраль
//  █        th 0 recieved Июнь
//  █        th 0 recieved Июль
//  █        th 0 recieved Август
//  █        th 0 recieved Сентябрь
//  █        th 0 recieved Октябрь
//  █        th 0 recieved Ноябрь
//  █        th 0 recieved Декабрь
//    █      th 1 recieved Май
//  █        === 0 finished
//    █      === 1 finished
//      █    th 2 recieved Апрель
//      █    === 2 finished
