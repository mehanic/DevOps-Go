package main

import "fmt"

func BinarySearch(array []int, target int) int {
	left, right := 0, len(array)-1

	for left <= right {
		middle := (left + right) / 2
		potentialMatch := array[middle]

		if target == potentialMatch {
			return middle
		} else if target < potentialMatch {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}

func main() {
	array := []int{11, 3, 5, 12, 9, 11, 52, 15, 104}
	targets := []int{7, 15, 2, 11}
	for _, target := range targets {
		result := BinarySearch(array, target)
		if result != -1 {
			fmt.Printf("Target %d found at index %d\n", target, result)
		} else {
			fmt.Printf("Target %d not found in the array\n", target)
		}
	}
}
