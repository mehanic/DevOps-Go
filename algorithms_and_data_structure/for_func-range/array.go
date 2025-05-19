package main

import "fmt"

func ListCounter(list []int) int {
	counter := 0

	for _, number := range list {
		counter += number
	}
	return counter
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := ListCounter(numbers)
	fmt.Println("Sum:", sum) // Output: Sum: 15
}
