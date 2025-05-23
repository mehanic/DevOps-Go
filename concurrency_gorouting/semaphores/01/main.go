package main

import (
	"fmt"
)

func main() {

	c := make(chan string)
	done := make(chan bool) 

	putInt := func(name string) {
		// wg.Add(1) 
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
		}
		done <- true 
	}

	go putInt("foo")
	go putInt("bar")

	go func() {
		<-done
		<-done 
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

}

// foo 0
// foo 1
// foo 2
// foo 3
// foo 4
// foo 5
// bar 0
// foo 6
// bar 1
// foo 7
// bar 2
// bar 3
// bar 4
// bar 5
// bar 6
// bar 7
// bar 8
// bar 9
// foo 8
// foo 9
