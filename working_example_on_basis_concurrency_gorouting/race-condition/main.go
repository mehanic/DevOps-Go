package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

var counter int

func main() {

	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// Bar 0 Counter: 1
// Foo 0 Counter: 1
// Bar 1 Counter: 2
// Foo 1 Counter: 2
// Bar 2 Counter: 3
// Bar 3 Counter: 4
// Bar 4 Counter: 5
// Foo 2 Counter: 3
// Bar 5 Counter: 6
// Foo 3 Counter: 4
