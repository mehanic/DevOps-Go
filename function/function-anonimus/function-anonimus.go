package main

import "fmt"

func main() {
	// Define an anonymous function and assign it to a variable
	add := func(x, y int) int {
		return x + y
	}
	
	// Use the anonymous function
	result := add(2, 3)
	fmt.Println(result)
}

