package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 50; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

// Bar:  0
// Foo:  0
// Foo:  1
// Foo:  2
// Bar:  1
// Foo:  3
// Foo:  4
// Bar:  2
// Foo:  5
// Foo:  6
