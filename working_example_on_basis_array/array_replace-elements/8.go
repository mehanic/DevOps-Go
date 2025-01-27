package main

import "fmt"

func replaceElements(a []int) []int {
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			a[i] = 0
		}
	}
	return a
}

// Logic:
// It iterates through each element of the slice using a for loop.
// For each element, it checks if the number is even (i.e., a[i] % 2 == 0).
// If the number is even, it replaces that element with 0.
// Return Value: The modified slice is returned.
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := replaceElements(arr)
	fmt.Println(arr)
	fmt.Println(arr2)
}

// [1 0 3 0 5 0 7 0 9 0]
// [1 0 3 0 5 0 7 0 9 0]
