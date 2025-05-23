package main

import (
	"fmt"
)

func main() {
	x := 99

	// Calculate cube, quotient, and remainder
	cube := x * x * x
	quotient := x / 5
	remainder := x % 5

	// Print the results
	fmt.Printf("The cube of %d is %d.\n", x, cube)
	fmt.Printf("The quotient of %d divided by 5 is %d.\n", x, quotient)
	fmt.Printf("The remainder of %d divided by 5 is %d.\n", x, remainder)

	// Calculate y using parentheses for operator precedence
	y := 2 * (quotient*5 + remainder)
	fmt.Printf("Two times %d is %d.\n", x, y)
}

