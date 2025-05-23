package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}

// Hello from goroutine 33
// Hello from goroutine 33
// Hello from goroutine 33
// Hello from goroutine 33
// Hello from goroutine 6
// Hello from goroutine 21
// Hello from goroutine 45
// Hello from goroutine 45
// Hello from goroutine 45
// Hello from goroutine 45
// Hello from goroutine 45
// Hello from goroutine 45
