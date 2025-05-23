package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sayHello(ch) //  без go будет deadlock
	s := <-ch
	fmt.Println(s)

}

func say2(mes string) {
	fmt.Println(mes)
}

func sayHello(exit chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		say2("Hello")
	}
	exit <- "exit"
}

// Hello
// Hello
// Hello
// Hello
// Hello
// exit
