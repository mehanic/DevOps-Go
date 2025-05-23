package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := 2
	c := make(chan string)
	done := make(chan bool) 

	putInt := func(name string) {
		// wg.Add(1) 
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
		}
		done <- true 
	}

	for i := 0; i < n; i++ {
		go putInt(strconv.Itoa(i))
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}

// 1 0
// 1 1
// 1 2
// 1 3
// 1 4
// 1 5
// 1 6
// 1 7
// 1 8
// 1 9
// 0 0
// 0 1
// 0 2
// 0 3
// 0 4
// 0 5
// 0 6
// 0 7
// 0 8
// 0 9
