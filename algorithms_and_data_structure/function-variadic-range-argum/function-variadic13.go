package main

import "fmt"

func sum2(num1 int, nums ...int) int {
	total := num1
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	//fmt.Println(sum2()) // gives error

	fmt.Println(sum2(0))
	fmt.Println(sum2(1, 2, 3, 4, 5))
}
