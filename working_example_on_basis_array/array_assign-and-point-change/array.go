package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 100
	fmt.Println(x)
	fmt.Println(len(x))
}

// [0 0 0 0 0]
// [0 0 0 100 0]
// 5
