package main

import (
	"fmt"
	"time"
)


func task1(result chan string) {
	time.Sleep(time.Second * 2)
	
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finised!")

}

func main() {
	result := make(chan string)

	go task1(result)
	go task2()

	fmt.Println(<-result)

	time.Sleep(time.Second * 3)

}

// task2 finised!
// task1 result
