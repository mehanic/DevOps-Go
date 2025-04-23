package main

import (
	"fmt"
	"sort"
)


//2. Sorting

func twoSum(nums []int, target int) []int {
	A := make([][2]int, len(nums))
	for i, num := range nums {
		A[i] = [2]int{num, i}
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i][0] < A[j][0]
	})

	i, j := 0, len(nums)-1
	for i < j {
		cur := A[i][0] + A[j][0]
		if cur == target {
			if A[i][1] < A[j][1] {
				return []int{A[i][1], A[j][1]}
			} else {
				return []int{A[j][1], A[i][1]}
			}
		} else if cur < target {
			i++
		} else {
			j--
		}
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
