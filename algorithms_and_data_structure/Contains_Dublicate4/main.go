// 1. Brute Force
package main

import (
	"fmt"
)

func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		seen[num] = struct{}{}
	}
	return len(seen) < len(nums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(hasDuplicate(nums))

	//nums1 := []int{1, 2, 3, 4, 3}
	fmt.Println(hasDuplicate([]int{1, 2, 3, 4, 3})) // true
	//nums2 := []int{5, 5, 5, 5, 5}
	fmt.Println(hasDuplicate([]int{5, 5, 5, 5, 5})) // true

	//nums3 := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(hasDuplicate([]int{1, 2, 3, 1, 2, 3})) // true
}
