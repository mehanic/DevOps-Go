package main

import "fmt"

// Declare func type that will always receive two int parameters and return an int.
type operation func(num1 int, num2 int) int

// Declare a function that takes the first argument of type operation, second and third parameters are arguments to first one.
func calculate(op operation, num1 int, num2 int) int {
	//Process passed operation and return the result.
	return op(num1, num2)
}

func main() {
	num1, num2 := 10, 5

	//Passing add function as the operation parameter
	result := calculate(add, num1, num2)
	fmt.Println(result) // This will print 15

	//Passing sub function as the operation parameter
	result = calculate(sub, num1, num2)
	fmt.Println(result) // This will print 2

	//Passing mul function as the operation parameter
	result = calculate(mul, num1, num2)
	fmt.Println(result) // This will print 50

	//Passing div function as the operation parameter
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

