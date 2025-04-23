package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i, n := range nums {
		diff := target - n
		if j, found := prevMap[diff]; found {
			return []int{j, i}
		}
		prevMap[n] = i
	}
	return []int{}
}

func main() {
	// Test cases
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))   // [0, 1] -> 2 + 7 = 9
	fmt.Println(twoSum([]int{3, 2, 4}, 6))        // [1, 2] -> 2 + 4 = 6
	fmt.Println(twoSum([]int{1, 5, 3, 8}, 9))     // [0, 3] -> 1 + 8 = 9
	fmt.Println(twoSum([]int{3, 3}, 6))           // [0, 1] -> 3 + 3 = 6
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 10)) // [] (no solution)
}
