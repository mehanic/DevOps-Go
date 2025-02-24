package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, th int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock()
				counters[th*10+j]++
				mu.Unlock()
			}
		}(counters, i, mu)
	}
	fmt.Scanln()
	fmt.Println("counters result", counters)
}

// ─λ go run race_2.go                                                                                                                                                 0 (2.453s) < 19:42:10

// counters result map[0:1 1:1 2:1 3:1 4:1 10:1 11:1 12:1 13:1 14:1 20:1 21:1 22:1 23:1 24:1 30:1 31:1 32:1 33:1 34:1 40:1 41:1 42:1 43:1 44:1]
// ╭─
