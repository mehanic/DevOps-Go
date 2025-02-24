package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar: ", i)
	}
	wg.Done()
}

// Bar:  0
// Bar:  1
// Bar:  2
// Bar:  3
// Bar:  4
// Bar:  5
// Bar:  6
// Bar:  7
// Bar:  8
// Bar:  9
// Foo:  0
// Foo:  1
// Foo:  2
// Foo:  3
// Foo:  4
// Foo:  5
// Foo:  6
// Foo:  7
// Foo:  8
// Foo:  9
