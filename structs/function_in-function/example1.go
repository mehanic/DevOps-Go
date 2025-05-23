package main

import "fmt"

// declare function type
type Operation func(balance, amount int) int

// function whose return type is an Operation (which is also a function)
func createOperation(typeOp string) Operation {
	if typeOp == "add" {
		return func(balance, amount int) int { return balance + amount }
	} else if typeOp == "subtract" {
		return func(balance, amount int) int { return balance - amount }
	} else {
		return func(balance, amount int) int { return balance * amount }
	}
}

func main() {
	add := createOperation("add")
	result := add(10, 15)
	fmt.Println("result:", result)
}

// result: 25
