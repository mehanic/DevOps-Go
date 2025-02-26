package main

import "fmt"

// Function to append an element to a slice and return the updated slice
func doAppend(sl []int) []int {
	return append(sl, 100)
}

func main() {
	// Initialize the slice
	x := []int{1, 2, 3}

	// Append 100 to the slice using the doAppend function
	x = doAppend(x)

	// Print the updated slice
	fmt.Println("outside: ", x) // Output: outside: [1 2 3 100]

	// Iterate over the slice and print the value without using the index
	for _, val := range x {
		fmt.Printf("slice entry: %d\n", val) // Output: slice entry: 1, slice entry: 2, slice entry: 3, slice entry: 100
	}
}

// outside:  [1 2 3 100]
// slice entry: 1
// slice entry: 2
// slice entry: 3
// slice entry: 100
