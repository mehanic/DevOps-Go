package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go say("dauren") // этот поток выполняется
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	time.Sleep(2 * time.Second) //здесь засыпает основной поток
	// и тут просыпается go say

}

func say(mes string) {
	time.Sleep(1 * time.Second) // здесь начинает спать
	fmt.Println(mes)
}

// ╭─mehanic at SkyNet in ⌁/algorithm_and_data_structure/example/go_lessons2/9/goroutine-basic
// ╰─λ go run goroutine.go                                                                                                                                              0 (0.002s) < 13:41:59
// start
// 1
// 2
// 3
// 4
// 5
// dauren
