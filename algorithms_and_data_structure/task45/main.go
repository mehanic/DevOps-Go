package main

import (
	"fmt"
	"sort"
)

// 1. Sorting
func missingNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums) // Sort the array
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i // Return the missing number
		}
	}
	return n // If no mismatch is found, return n
}

func main() {
	// Example 1: Missing number is 0
	nums1 := []int{1, 2, 3}
	fmt.Println("Missing number in [1, 2, 3]:", missingNumber(nums1)) // Output: 0

	// Example 2: Missing number is 1
	nums2 := []int{0, 2}
	fmt.Println("Missing number in [0, 2]:", missingNumber(nums2)) // Output: 1

	// Example 3: Missing number is 2
	nums3 := []int{0, 1}
	fmt.Println("Missing number in [0, 1]:", missingNumber(nums3)) // Output: 2

	// Example 4: Missing number is 2
	nums4 := []int{0, 1, 3, 4}
	fmt.Println("Missing number in [0, 1, 3, 4]:", missingNumber(nums4)) // Output: 2

	// Example 5: Missing number is 4
	nums5 := []int{0, 1, 2, 3, 5}
	fmt.Println("Missing number in [0, 1, 2, 3, 5]:", missingNumber(nums5)) // Output: 4
}

//2. Hash Set

func missingNumber1(nums []int) int {
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	n := len(nums)
	for i := 0; i <= n; i++ {
		if _, exists := numSet[i]; !exists {
			return i
		}
	}
	return -1
}

//3. Bitwise XOR

func missingNumber2(nums []int) int {
	n := len(nums)
	xorr := n
	for i := 0; i < n; i++ {
		xorr ^= i ^ nums[i]
	}
	return xorr
}

//4. Math

func missingNumber3(nums []int) int {
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res += i - nums[i]
	}
	return res
}
