package main

import "fmt"

// Function that simulates `myfunction1`
func myFunction1() {
	x := 60 // This is a local variable
	fmt.Println("Welcome to functions")
	fmt.Printf("x value from fun1: %d\n", x)
}

// Function that simulates `myfunction2`
func myFunction2(y int) { // Parameter
	fmt.Println("Thank you!!")
	fmt.Printf("x value from fun2: %d\n", y)
}

func main() {
	x := 10 // Local variable
	myFunction1()
	myFunction2(x) // Argument
}

// Welcome to functions
// x value from fun1: 60
// Thank you!!
// x value from fun2: 10
