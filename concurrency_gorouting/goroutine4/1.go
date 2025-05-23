package main

import (
	"fmt"
	"time"
)

func main() {
	go printData()
	time.Sleep(5 * time.Second)
	fmt.Println("good")
}

func printData() {
	fmt.Println("this is goroutine mesasge")
	time.Sleep(100 * time.Second)
}

// this is goroutine mesasge
// good
