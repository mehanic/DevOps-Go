package main

import (
	"fmt"
	"sort"
)

//2. Sorting

func hasDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
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
