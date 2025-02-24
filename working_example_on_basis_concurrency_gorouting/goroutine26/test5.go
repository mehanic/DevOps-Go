package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("hello World")
	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("hello bello world")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}
	time.Sleep(2 * time.Second)
	fmt.Scanln()
}

// ─λ go run test5.go                                                                       0 (0.001s) < 15:45:37
// hello World
// timer2 stopped
