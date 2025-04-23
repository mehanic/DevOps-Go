package main

import "fmt"

// O(log(n)) time | O(log(n)) space
func BinarySearch(array []int, target int) int {
	return helper(array, target, 0, len(array)-1)
}

// Recursive helper function
func helper(array []int, target, left, right int) int {
	// Base case: If the range is invalid, target is not found
	if left > right {
		return -1
	}

	// Find middle index
	middle := (left + right) / 2
	potentialMatch := array[middle]

	// Check if middle element is the target
	if target == potentialMatch {
		return middle
	} else if target < potentialMatch {
		// If target is smaller, search the left half
		return helper(array, target, left, middle-1)
	}

	// If target is larger, search the right half
	return helper(array, target, middle+1, right)
}

func main() {
	// Example sorted array
	array := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}

	// Test cases
	targets := []int{7, 15, 2, 11}
	for _, target := range targets {
		result := BinarySearch(array, target)
		if result != -1 {
			fmt.Printf("Target %d found at index %d\n", target, result)
		} else {
			fmt.Printf("Target %d not found in the array\n", target)
		}
	}

	fmt.Println("----------------")
	main1()
}

// O(log(n)) time | O(1) space
func BinarySearch1(array []int, target int) int {
	left, right := 0, len(array)-1

	for left <= right {
		middle := (left + right) / 2
		potentialMatch := array[middle]

		if target == potentialMatch {
			return middle // Found target, return its index
		} else if target < potentialMatch {
			right = middle - 1 // Move to the left half
		} else {
			left = middle + 1 // Move to the right half
		}
	}

	return -1 // Target not found
}

func main1() {
	// Example sorted array
	array := []int{11, 3, 5, 12, 9, 11, 52, 15, 104}

	// Test cases
	targets := []int{7, 15, 2, 11}
	for _, target := range targets {
		result := BinarySearch1(array, target)
		if result != -1 {
			fmt.Printf("Target %d found at index %d\n", target, result)
		} else {
			fmt.Printf("Target %d not found in the array\n", target)
		}
	}
}
