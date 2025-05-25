package main

import "fmt"

func helper(array []int, target, left, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2
	potentialMatch := array[middle]

	if target == potentialMatch {
		return middle
	} else if target < potentialMatch {
		return helper(array, target, left, middle-1)
	}
	return helper(array, target, middle+1, right)
}

func main() {
	array := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	targets := []int{7, 15, 2, 11}

	for _, target := range targets {
		index := helper(array, target, 0, len(array)-1)
		if index != -1 {
			fmt.Printf("Target %d found at index %d\n", target, index)
		} else {
			fmt.Printf("Target %d not found\n", target)
		}
	}
}