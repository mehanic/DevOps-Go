package main

import (
	"fmt"
	"time"
)

type genInt64 int64

func (g *genInt64) Next() int64 {
	*g++
	return int64(*g)
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

// ╭─mehanic at SkyNet in ⌁/algorithm_and_data_structure/example/generator/generator-naive
// ╰─λ go run naivegen.go                                                                                                                                               0 (0.002s) < 22:54:45
// 0 1
// 3 3
// 1 4
// 2 5
// 9 6
// 8 7
// 5 8
// 4 9
// 6 10
// 7 2
// 10 15
// 16 16
// 19 20
