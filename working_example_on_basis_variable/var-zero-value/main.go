package main

import "fmt"

// 在golang, 變數的預設值稱為"zero-value"

func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}

// 0

// 0
// false
