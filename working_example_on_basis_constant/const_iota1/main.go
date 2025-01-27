package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
)

// defining a second iota will start at 0
const (
	d = iota // 0
	e        // 1
	f        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

// 0
// 1
// 2
// 0
// 1
// 2
