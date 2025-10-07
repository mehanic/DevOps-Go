package main

import "fmt"

func moveZeroes(nums []int) {
	insertPos := 0

	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}

	for i := insertPos; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	// Test Case 1
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(nums)
	moveZeroes(nums)
}
