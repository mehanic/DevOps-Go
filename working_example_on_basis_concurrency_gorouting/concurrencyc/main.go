package main

import (
	"fmt"
	"time"
)

func heavyTask(n int, c chan string) {

	c <- fmt.Sprintf("task %d has began \n", n)

	time.Sleep(time.Duration(n) * time.Second)

	c <- fmt.Sprintf("%d has completed \n", n)
}

func main() {
	c := make(chan string)

	for i := 0; i < 5; i++ {
		go heavyTask(i, c)
	}

	result := make([]string, 10)

	for i := range result {
		result[i] = <-c
		fmt.Println(result[i])

	}

}

// ─λ go run main.go                                                                        0 (0.001s) < 13:28:02
// task 4 has began

// task 0 has began

// 0 has completed

// task 1 has began

// task 2 has began

// task 3 has began

// 1 has completed

// 2 has completed

// 3 has completed

// 4 has completed
