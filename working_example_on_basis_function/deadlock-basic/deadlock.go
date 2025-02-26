package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sayHello2(ch) //  без go будет deadlock

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch) //deadlock

	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(<-ch)
}

func helo(mes string) {
	fmt.Println(mes)
}

func sayHello2(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
	close(exit) // закрытие канала
	//exit <- 1 в закрытый канал нельзя ничего послать

}
