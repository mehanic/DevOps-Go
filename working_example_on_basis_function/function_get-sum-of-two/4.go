package main

import "fmt"

func getSumOfTwo(a, b int) int {
	c := a + b
	return c
}

func main() {
	d := getSumOfTwo(3, 4)
	l := getSumOfTwo(d, 10)
	fmt.Println(l)
}
