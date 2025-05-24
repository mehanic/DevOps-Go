package main

import "fmt"

func insertAtIndex(array []int, value int, index int) []int {
	if index < 0 || index > len(array) {
		fmt.Println("Index out of range")
		return array
	}

	// Create a new slice with space for the new value
	newArray := append(array[:index], append([]int{value}, array[index:]...)...)

	return newArray
}

func main() {
	// Given array
	A := []int{22, 68, 67, 34, 12, 3, 8, 9, 0, 34, 81, 64, 46}

	// Insert 88 at index 8
	A = insertAtIndex(A, 99, 8)

	// Print the updated array
	fmt.Println("Updated Array:", A)
}
