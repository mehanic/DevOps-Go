package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan string)

	putInt := func(name string) {
		// wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", name, i)
		}
		wg.Done()
	}

	wg.Add(2)
	go putInt("foo")
	go putInt("bar")

	go func() {
		wg.Wait()
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
// bar 0
// bar 1
// foo 4
// foo 5
// foo 6
// foo 7
// foo 8
// foo 9
// bar 2
// bar 3
// bar 4
// bar 5
// bar 6
// bar 7
// bar 8
// bar 9
