package main

import "fmt"

func main() {

	// passing multiple argument
	fmt.Println(sum(1, 2, 3, 4, 5))

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(xi...)) // passing all value in slice xi, variadic argument

	// fmt.Println(sum(xi)) // error, xi is a slice
}

// variadic parameter (...type)
func sum(xi ...int) int {
	result := 0
	for _, v := range xi {
		result += v
	}
	return result
}
