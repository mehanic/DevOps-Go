package main

import "fmt"

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 30; i++ {
			c <- i
		}
		close(c)
	}()

	n := 5
	for i := 0; i < n; i++ {
		go func() {
			for i := range c {
				fmt.Println(i)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}

}

// 3
// 0
// 1
// 2
// 4
// 9
// 10
// 11
// 12
// 13
// 14
// 15
// 16
// 17
// 18
// 19
// 20
// 21
// 22
// 23
// 24
// 25
// 26
// 27
// 28
// 29
// 7
// 5
// 8
// 6
