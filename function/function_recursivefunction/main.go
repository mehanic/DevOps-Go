package main

import (
	"fmt"
)

func main() {
	x := factorial(4)
	fmt.Println(x)
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

//function calls itself is called recursion and factorial is recursive function eg
// ie factorial of 4 is 4*3*2*1
