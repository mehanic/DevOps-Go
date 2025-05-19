package main

import "fmt"

/*
MathOperation : The Math Operation function accepts three parameters, num1, num2,and the third parameter

	is a function that performs one of the basic operations on num1 and num2.
*/
func MathOperation(num1 int, num2 int, calculate func(int, int) int) int {
	return calculate(num1, num2)
}

func main() {
	num1, num2 := 10, 5

	//Passing add function as the third parameter
	operation := MathOperation(num1, num2, add)
	fmt.Println(operation) //This will print 15

	//Passing sub function as the third parameter
	operation = MathOperation(num1, num2, sub)
	fmt.Println(operation) //This will print 5

	//Passing mul function as the third parameter
	operation = MathOperation(num1, num2, mul)
	fmt.Println(operation) //This will print 50

	//Passing div function as the third parameter
	operation = MathOperation(num1, num2, div)
	fmt.Println(operation) //This will print 2
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

// 15
// 5
// 50
// 2
