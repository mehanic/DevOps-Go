package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // define an unbuffured channel. Channel without specifying allocation
	//c := make(chan int, 10) // define an buffured channel.

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // put the value of i into the channel.
			//Here the code stops until something receives the value from channel
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // take the number of the channel and print that value
		}
	}()

	time.Sleep(time.Second)

}

// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
