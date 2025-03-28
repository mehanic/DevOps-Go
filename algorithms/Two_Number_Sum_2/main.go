package main

import (
	"fmt"
	"sort"
)

// O(nlog(n)) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}


func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 9
	result := TwoNumberSum(array, target)
	fmt.Println(result, "this is result of array := []int{3, 5, -4, 8, 11, 1, -1, 6}") 
}
