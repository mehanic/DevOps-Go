package main

import (
	"fmt"
)

var (
	cha    = make(chan int)
	global = make(chan int)
)

func Set() {
	fmt.Println("set run")
	cha <- 3
}

func Get() {
	fmt.Println("get run")
	fmt.Println(<-cha)
	cha <- 4
	fmt.Println("end of get")
	global <- 5
}

func main() {
	go Set()
	go Get()
	<-global
}

// get run
// set run
// 3
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
// 	/home/mehanic/structure/concurrency_i_goroutines/175/5.go:28 +0x34

// goroutine 19 [chan send]:
// main.Get()
// 	/home/mehanic/structure/concurrency_i_goroutines/175/5.go:20 +0xbe
// created by main.main in goroutine 1
// 	/home/mehanic/structure/concurrency_i_goroutines/175/5.go:27 +0x26
// exit status 2
