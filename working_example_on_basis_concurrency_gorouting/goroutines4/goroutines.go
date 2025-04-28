package main

import (
	"fmt"
	"time"
)

func printMessage(message string, wait time.Duration) {
	time.Sleep(wait)
	fmt.Println(message)
}

func main() {

	go printMessage("start", 3*time.Second)

	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("time:", i)
	}

	go printMessage("end", 3*time.Second)

	time.Sleep(4 * time.Second)
}

// time: 0
// time: 1
// time: 2
// start
// time: 3
// time: 4
// time: 5
// end
