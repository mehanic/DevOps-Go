package main

import (
	"fmt"
	"math"
)

// O(n) time | O(1) space
func FindThreeLargestNumbers(array []int) []int {
	three := []int{math.MinInt32, math.MinInt32, math.MinInt32} // Initialize with smallest possible integers
	for _, num := range array {
		updateLargest(three, num)
	}
	return three
}

func updateLargest(three []int, num int) {
	if num > three[2] {
		shiftAndUpdate(three, num, 2) // Update largest
	} else if num > three[1] {
		shiftAndUpdate(three, num, 1) // Update second largest
	} else if num > three[0] {
		shiftAndUpdate(three, num, 0) // Update third largest
	}
}

func shiftAndUpdate(array []int, num int, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			array[i] = num
		} else {
			array[i] = array[i+1]
		}
	}
}

func main() {
	// Example 1
	array1 := []int{10, 4, 3, 50, 23, 90}
	result1 := FindThreeLargestNumbers(array1)
	fmt.Println("Three largest numbers:", result1) // Expected Output: [23, 50, 90]

	// Example 2
	array2 := []int{-5, -10, -1, -8, -3, -7}
	result2 := FindThreeLargestNumbers(array2)
	fmt.Println("Three largest numbers:", result2) // Expected Output: [-5, -3, -1]

	// Example 3
	array3 := []int{100, 200, 300, 400, 500}
	result3 := FindThreeLargestNumbers(array3)
	fmt.Println("Three largest numbers:", result3) // Expected Output: [300, 400, 500]
}
