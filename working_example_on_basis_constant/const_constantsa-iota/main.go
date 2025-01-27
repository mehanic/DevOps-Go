package main

import "fmt"

// untyped constant
const a = 10

// typed constant
const b string = "hello"

// const arr []int = []int{0, 1, 2} // error, array can not be constant

// assign multiple constant
const (
	c = 10
	d = "oh"
	e = true
)

// iota: multiple-initialize
const (
	_ = iota       // 0
	f              // 1
	g              // 2
	h              // 3
	i = iota*2 + 1 // 4*2+1 = 9
)

func main() {

	// a = "else" // error, constant is unchanged value
	fmt.Println(a)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	const q = "hey"
	fmt.Println(q)

}
