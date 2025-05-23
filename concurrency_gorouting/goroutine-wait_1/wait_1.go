package main

import (
	"log"
	"time"
)

func main() {
	result := make(chan string)
	go func(out chan<- string) {
		time.Sleep(1 * time.Second)
		log.Println("async operation ready, return result")
		out <- "success"
	}(result)

	time.Sleep(2 * time.Second)
	log.Println("some userful work")

	opStatus := <-result
	log.Println("main goroutine:", opStatus)
}

// ╰─λ go run wait_1.go                                                                                                                                                 0 (0.001s) < 19:45:57
// 2024/09/21 19:46:00 async operation ready, return result
// 2024/09/21 19:46:01 some userful work
// 2024/09/21 19:46:01 main goroutine: success
