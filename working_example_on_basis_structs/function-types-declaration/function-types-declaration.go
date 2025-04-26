package main

import "fmt"

type operation func(num1 int, num2 int) int

func calculate(op operation, num1 int, num2 int) int {
	return op(num1, num2)
}

func main() {
	num1, num2 := 10, 5

	result := calculate(add, num1, num2)
	fmt.Println(result) // This will print 15

	result = calculate(sub, num1, num2)
	fmt.Println(result) // This will print 2

	result = calculate(mul, num1, num2)
	fmt.Println(result) // This will print 50

	result = calculate(div, num1, num2)
	fmt.Println(result) // This will print 5
}

func sub(num1, num2 int) int {
	return num1 - num2
}

func add(num1, num2 int) int {
	return num1 + num2
}

func mul(num1, num2 int) int {
	return num1 * num2
}

func div(num1, num2 int) int {
	return num1 / num2
}
