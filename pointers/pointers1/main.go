package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var pointer *int
	// fmt.Println("Value of pointer is ", pointer)

	myNumber := 23

	var pointer = &myNumber

	fmt.Println("Value of actual pointer is ", pointer)
	fmt.Println("Value of actual pointer is ", *pointer)

	*pointer = *pointer * 2
	fmt.Println("New value is: ", myNumber)
}
