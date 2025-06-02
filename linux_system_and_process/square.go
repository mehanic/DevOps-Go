package main

import (
	"fmt"
)

func SquareSum(numbers []int) int {
	// idx := 1
	// val := 0
	sum := 0
	// result = 1
	// for _, num := range numbers {
	// 	fmt.Println(num)
	// 	// result = (num * num) + num
	// }
	// return (num * num) + num
	for idx, val := range numbers {
		fmt.Println(idx, val)
		sum += (val * val)
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(SquareSum(numbers))
}
