package main

import "fmt"

// out of func scope is package scope

func main() { // func scope
	x := 30
	fmt.Println(x)
	// scope is surround by curly braces
	{
		fmt.Println(x) // inner scope can see outer scope

		y := 40
		fmt.Println(y)
	}
	// fmt.Println(y) // error, outer scope can't see inner scope

	x = 0
	increment := func() int {
		x++ // inner scope can see outer scope, so x is remembered
		return x
	}
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}
