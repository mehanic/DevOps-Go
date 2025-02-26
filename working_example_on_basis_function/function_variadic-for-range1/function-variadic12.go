package main

import "fmt"

func sum1(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum1())     // with zero (empty) argument
	fmt.Println(sum1(1, 2)) // more than zero argument
}
