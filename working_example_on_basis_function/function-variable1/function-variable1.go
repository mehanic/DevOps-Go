package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {

	var f func(int, int) int = add
	fmt.Println(f(3, 4))

	var x = f(4, 5) // 9
	fmt.Println(x)
}

