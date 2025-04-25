package main

import (
	"fmt"
)

func main() {
	f1 := func() {
		fmt.Println("Function expression")
	}
	f1()
	// func()
	fmt.Printf("%T\n", f1)

	f2 := func(x int, y int) {
		fmt.Println("Sum of x and y is", x+y)
	}
	f2(2, 3)
}

// Function expression
// func()
// Sum of x and y is 5
