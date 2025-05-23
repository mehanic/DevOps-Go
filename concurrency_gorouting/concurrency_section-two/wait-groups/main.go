package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha", "beta", "delta", "gamma", "pi", "zeta", "eta", "theta", "epsilon",
	}

	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
	wg.Add(1)
	printSomething("Thi is the 2nd thing to be printed", &wg)
}

// ─λ go run main.go                                                                        0 (0.001s) < 12:57:51
// 8: epsilon
// 4: pi
// 5: zeta
// 6: eta
// 7: theta
// 0: alpha
// 1: beta
// 2: delta
// 3: gamma
// Thi is the 2nd thing to be printed
