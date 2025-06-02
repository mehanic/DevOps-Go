package main

import "fmt"

func PositiveSum(numbers []int) int {
	res := 0
	for _, value := range numbers {
		if value > 0 {
			// x := int(value)
			// fmt.Println(value)
			res += value

		}

	}
	return res

}

func main() {

	numbers := []int{-2, -1, 0, 1, 2, -9, 5, 7}
	fmt.Println(PositiveSum(numbers))

}
