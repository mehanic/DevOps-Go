package main

import "fmt"

func main() {
	factorial := factorialRecursive(10)
	fmt.Println(factorial) // 3628800
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}