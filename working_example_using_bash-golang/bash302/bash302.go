package main

import (
	"fmt"
)

func main() {
	// Declare integers
	num1 := 2
	num2 := 1

	// Compare integers and branch accordingly
	if num1 == num2 {
		fmt.Println("Both Values are equal")
	} else if num1 > num2 {
		fmt.Println("NUM1 is greater than NUM2")
	} else {
		fmt.Println("NUM2 is greater than NUM1")
	}
}

