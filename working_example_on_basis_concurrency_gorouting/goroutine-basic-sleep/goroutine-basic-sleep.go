package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go say("dauren", ch) // этот поток выполняется
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println(<-ch)
	//time.Sleep(2 * time.Second) //здесь засыпает основной поток
	// и тут просыпается go say

}

func say(mes string, ch chan int) {
	time.Sleep(10 * time.Second) // здесь начинает спать
	fmt.Println(mes)
	ch <- 1
}

// ╰─λ go run g2.go                                                                                                                                                     0 (0.001s) < 02:03:24
// 1
// 2
// 3
// 4
// 5
// dauren
// 1
