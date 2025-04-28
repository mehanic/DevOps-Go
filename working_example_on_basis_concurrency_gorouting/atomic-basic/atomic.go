package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter uint64
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1) // n gorutin
		//k := i
		go func() {
			defer wg.Done()
			for g := 0; g < 100; g++ {
				mu.Lock()
				counter++
				mu.Unlock()
				//atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Print("done...", counter)

}

// done...1000⏎
