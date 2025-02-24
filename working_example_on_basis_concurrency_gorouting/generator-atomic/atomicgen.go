package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type genInt64 int64

func (g *genInt64) Next() int64 {
	return atomic.AddInt64((*int64)(g), 1)
}

func main() {
	var g genInt64
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i, g.Next())
		}(i)
	}
	time.Sleep(time.Second)
}

// ╰─λ go run main.go                                                                                                               0 (0.000s) < 13:44:58
// Received:  Hello!
// Received:  Salut!
