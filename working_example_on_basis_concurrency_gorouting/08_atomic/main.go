package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var count int64
var wg sync.WaitGroup

func increment(name string) {
	for i := 1; i <= 20; i++ {
		time.Sleep(time.Duration(5 * time.Millisecond))
		atomic.AddInt64(&count, 1)
		
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()
	fmt.Println(count)
}

// 40
