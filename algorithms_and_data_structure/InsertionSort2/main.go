package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	length := len(arr)

	// Traverse through 1 to length of array
	for i := 1; i < length; i++ {
		j := i - 1

		// Swap elements that are not in the right position
		for j >= 0 && arr[j] > arr[j+1] {
			// Swap arr[j] and arr[j+1]
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Before Sorting:", arr)
	insertionSort(arr)
	fmt.Println("After Sorting:", arr)
}
