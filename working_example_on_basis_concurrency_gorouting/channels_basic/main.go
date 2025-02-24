package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch) // zero value of a channel is <nil>

	ch = make(chan int) //initializing a channel
	fmt.Println(ch)     //returns a pointer value

	c := make(chan int) //the above channel creation and initialization can be achieved with a one liner using this short declaration
	fmt.Println(c)

	// <- channel operator

	//SEND Operation
	c <- 10 //this sends the value 10 into the channel

	//RECEIVE operation
	num := <-c

	fmt.Println(num) //or
	fmt.Println(<-c)

	close(c) //this closes a channel

}

// ─λ go run main.go                                                                        0 (0.001s) < 13:31:09
// <nil>
// 0xc00008e070
// 0xc00008e0e0
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
// 	/home/mehanic/structure/concurrency_i_goroutines/concurrencyt7/channels_basic/main.go:18 +0xf5
// exit status 2
