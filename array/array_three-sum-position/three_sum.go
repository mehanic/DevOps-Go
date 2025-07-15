package main

import (
	"fmt"
	"sort"
)

// threeSum finds all unique triplets in the array which give the sum of zero.
func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s > 0 {
				r -= 1
			} else if s < 0 {
				l += 1
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			}
		}
	}
	return res
}

func main() {
	// Test array
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("Input array:", nums)

	// Call threeSum function
	result := threeSum(nums)

	// Display the result
	fmt.Println("Unique triplets that sum to zero:", result)
}

// Input array: [-1 0 1 2 -1 -4]
// Unique triplets that sum to zero: [[-1 -1 2] [-1 0 1]]
