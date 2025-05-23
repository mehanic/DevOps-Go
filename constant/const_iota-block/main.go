package main

import "fmt"

// saying iota once, rest will also be iota
const (
	a = iota // 0
	b        // 1
	c        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// 0
// 1
// 2
