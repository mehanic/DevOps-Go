package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	// 1. Reverse the entire array
	reverse(0, n-1) //[7, 6, 5, 4, 3, 2, 1]
	// 2. Reverse the first k elements
	reverse(0, k-1) //[5, 6, 7, 4, 3, 2, 1]
	// 3. Reverse the remaining n-k elements
	reverse(k, n-1) //[5, 6, 7, 1, 2, 3, 4]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)
	fmt.Println(nums) // Output: [5 6 7 1 2 3 4]
}
