package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	
	fmt.Println("number of cpu:", runtime.NumCPU())
	fmt.Println("previous gomaxprocs:", runtime.GOMAXPROCS(runtime.NumCPU()))
}

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

// number of cpu: 12
// previous gomaxprocs: 12
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
// Bar:  3
// Foo:  7
// Bar:  4
// Foo:  8
