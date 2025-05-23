package main

import "fmt"

var global = func() int {
	return 10
}

func simple_function(x int) int {
	func() {}()

	return x * x
}

func welcome(text string) {
	fmt.Println(text)
}

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func build(x func(int, int) int, i, k int) int {
	return x(i, k)
}

// div
// mult
// mod %

func main() {
	fmt.Println(build(add, 10, 20))
	fmt.Println(build(sub, 10, 20))
}
