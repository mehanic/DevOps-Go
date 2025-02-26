package main

import "fmt"

func factorial(n int) int {

	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {

	fmt.Println(factorial(4))
	fmt.Println(factorial(5))
}

