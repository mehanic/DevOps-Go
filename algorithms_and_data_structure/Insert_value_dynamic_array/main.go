package main

import "fmt"

// Function to insert a value into a dynamic array
func insertValue(arr []int, lenFilled *int, value, index int) []int {
	// Check if array is full, if yes, double its size
	if *lenFilled == len(arr) {
		newArr := make([]int, len(arr)*2) // Double the size
		copy(newArr, arr)                 // Copy old elements to new array
		arr = newArr
	}

	// Shift elements to the right to make space for the new value
	for j := *lenFilled - 1; j >= index; j-- {
		arr[j+1] = arr[j]
	}

	// Insert the new value
	arr[index] = value

	// Increase the filled length
	*lenFilled++

	return arr
}

func main() {
	// Initial array with some values and space for additional elements
	A := make([]int, 10) // Predefined size (can grow dynamically)
	A = []int{1, 2, 3, 4, 5, 6, 7, 9, 10} // Initially filled elements
	lenFilled := len(A)                   // Tracking filled size

	// Insert value 88 at index 7
	A = insertValue(A, &lenFilled, 88, 7)

	// Print the updated array
	fmt.Println("Updated Array:", A)
}
