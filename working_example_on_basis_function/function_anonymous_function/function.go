package main

import "fmt"

func main() {
	multiple := func(number1 int, number2 int) int {
		return number1 * number2
	}

	fmt.Println(multiple(2, 3)) // 6
}