package main

import "fmt"

func sum1(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func sum2(num1 int, nums ...int) int {
	total := num1
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	S := []int{1, 2, 3, 4, 5}
	fmt.Println(sum1(S...))

	// first argument should be normal
	fmt.Println(sum2(0, S...))
}

// 15
// 15

