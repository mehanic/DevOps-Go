package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-k-1) // Reverse the first part
	reverse(nums, n-k, n-1) // Reverse the second part
	reverse(nums, 0, n-1)   // Reverse the whole array
}

func reverse(slice []int, a, b int) {
	for a < b {
		slice[a], slice[b] = slice[b], slice[a]
		a += 1
		b -= 1
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Println("Original array:", nums)
	rotate(nums, k)
	fmt.Println("Array after rotation:", nums)
}

// Original array: [1 2 3 4 5 6 7]
// Array after rotation: [5 6 7 1 2 3 4]
