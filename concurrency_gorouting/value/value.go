package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		v  atomic.Value
		wg sync.WaitGroup
	)
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("load", v.Load())
			wg.Done()
		}(i)
		go func(i int) {
			v.Store(i)
			fmt.Println("store", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// store 9
// load 9
// store 0
// load 0
// store 1
// load 1
// store 2
// load 2
// store 6
// load 6
// store 7
// load 7
// store 8
// load 8
// store 4
// store 3
// load 3
// load 3
// store 5
// load 6
