package main

import (
	"fmt"
)

func main() {
	num1 := 2
	num2 := 3

	if num1 == num2 {
		fmt.Println("Both Values are equal")
	} else {
		fmt.Println("Values are not equal")
	}

	num3 := 4
	num4 := 4

	if num3 == num4 {
		fmt.Println("Both Values are equal")
	} else {
		fmt.Println("Values are not equal")
	}
}

