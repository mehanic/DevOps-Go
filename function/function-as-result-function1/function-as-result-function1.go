package main

import "fmt"

func add(x int, y int) int {

	return x + y
}
func multiply(x int, y int) int {

	return x * y
}
func action(n1 int, n2 int, operation func(int, int) int) {

	result := operation(n1, n2)
	fmt.Println(result)
}
func main() {

	action(10, 25, add)    // 35
	action(5, 6, multiply) // 30
}
