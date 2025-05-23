package main

import (
	"fmt"
)

func main() {
	// Define integer variables
	a := 12
	b := 24
	c := 78
	d := 24

	// Print the values of a, b, c, and d
	fmt.Printf("a = %d , b = %d , c = %d , d = %d\n", a, b, c, d)

	// Check if 'a' is greater than 'b'
	fmt.Print("Is a greater than b? ")
	if a > b {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}

	// Check if 'b' is equal to 'd'
	fmt.Print("Is b equal to d? ")
	if b == d {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}

	// Check if 'c' is not equal to 'd'
	fmt.Print("Is c not equal to d? ")
	if c != d {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}

