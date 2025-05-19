package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	var num1, num2 float32 = 5.6, 9.5
	fmt.Println(add(num1, num2))

	a, b := "hello", "my name Peter"
	fmt.Println(multiple(a, b))
}
