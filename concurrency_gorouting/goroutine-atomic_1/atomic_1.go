package main

import (
	"fmt"
	"time"
)

var totalOperations int32 = 0

func inc() {
	totalOperations++
}

func main() {
	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	// ождается 1000, но по факту будет меньше
	fmt.Println("total operation = ", totalOperations)
}

// ╰─λ go run atomic_1.go                                                                                                                                               0 (0.001s) < 19:26:35
// total operation =  994
