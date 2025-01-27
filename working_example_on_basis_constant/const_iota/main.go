package main

import "fmt"

// an iota here is very similar to an enum in C++ (increments value)
// I can also do:
// const (
//
//	A = iota
//	B
//	C
//
// )
// so I only have to type iota once
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// 0
// 1
// 2
