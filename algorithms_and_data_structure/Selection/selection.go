package main

import (
	"fmt"
)

func selectionSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		// Assume the current index as the minimum
		minimum := i

		for j := i + 1; j < n; j++ {
			if array[j] < array[minimum] {
				// Update the minimum index
				minimum = j
			}
		}

		// Swap the elements
		array[i], array[minimum] = array[minimum], array[i]
	}
	return array
}

func main() {
	array := []int{13, 4, 9, 5, 3, 16, 12}
	sortedArray := selectionSort(array)
	fmt.Println(sortedArray)
}
