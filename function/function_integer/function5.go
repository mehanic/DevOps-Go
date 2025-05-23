package main

import "fmt"

func main() {
	add1(1, 2, 3.4, 5.6, 1.2)
}
func add1(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

// x + y =  3
// a + b + c =  10.2