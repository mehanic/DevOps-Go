package main

import "fmt"

// O(n) time | O(n) space
func TwoNumberSum(array []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		} else {
			nums[num] = true
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


