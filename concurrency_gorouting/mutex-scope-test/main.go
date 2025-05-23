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
	wg.Add(3)
	go incrementor("Foo")
	go incrementor("Bar")
	go test()
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

const times int = 100

func incrementor(s string) {
	for i := 0; i < times; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}

func test() {
	for i := 0; i < times; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		counter = x
		fmt.Println("test", i, "Counter:", counter)
	}
	wg.Done()
}


// test 0 Counter: 1
// test 1 Counter: 2
// test 2 Counter: 3
// test 3 Counter: 4
// test 4 Counter: 5
// Foo 0 Counter: 6
// test 5 Counter: 6
// test 6 Counter: 7
// Bar 0 Counter: 8
// Bar 1 Counter: 9
// test 7 Counter: 8
// test 8 Counter: 9
// Bar 2 Counter: 10
// test 9 Counter: 10
