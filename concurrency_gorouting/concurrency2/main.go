package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // We have two goroutines

	go foo(&wg)
	go bar(&wg)

	wg.Wait() // Wait for both goroutines to finish
}

func foo(wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	for i := 0; i < 50; i++ {
		fmt.Println("Foo: ", i)
	}
}

func bar(wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	for i := 0; i < 50; i++ {
		fmt.Println("Bar: ", i)
	}
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
// Bar:  10
// Bar:  11
// Bar:  12
// Bar:  13
// Bar:  14
// Bar:  15
// Bar:  16
// Bar:  17
// Bar:  18
// Bar:  19
// Bar:  20
// Bar:  21
// Bar:  22
// Bar:  23
// Bar:  24
// Bar:  25
// Bar:  26
// Bar:  27
// Bar:  28
// Bar:  29
// Bar:  30
// Bar:  31
// Bar:  32
// Bar:  33
// Bar:  34
// Bar:  35
// Bar:  36
// Bar:  37
// Bar:  38
// Bar:  39
// Bar:  40
// Bar:  41
// Bar:  42
// Bar:  43
// Bar:  44
// Bar:  45
// Bar:  46
// Bar:  47
// Bar:  48
// Bar:  49
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
// Foo:  10
// Foo:  11
// Foo:  12
// Foo:  13
// Foo:  14
// Foo:  15
// Foo:  16
// Foo:  17
// Foo:  18
// Foo:  19
// Foo:  20
// Foo:  21
// Foo:  22
// Foo:  23
// Foo:  24
// Foo:  25
// Foo:  26
// Foo:  27
// Foo:  28
// Foo:  29
// Foo:  30
// Foo:  31
// Foo:  32
// Foo:  33
// Foo:  34
// Foo:  35
// Foo:  36
// Foo:  37
// Foo:  38
// Foo:  39
// Foo:  40
// Foo:  41
// Foo:  42
// Foo:  43
// Foo:  44
// Foo:  45
// Foo:  46
// Foo:  47
// Foo:  48
// Foo:  49
