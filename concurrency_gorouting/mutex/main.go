package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

// Foo 0 Counter: 1
// Bar 0 Counter: 2
// Bar 1 Counter: 3
// Foo 1 Counter: 4
// Bar 2 Counter: 5
// Foo 2 Counter: 6
// Bar 3 Counter: 7
// Foo 3 Counter: 8
// Bar 4 Counter: 9
