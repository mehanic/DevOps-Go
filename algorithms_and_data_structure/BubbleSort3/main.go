// Code for Bubble Sort in Go
package main

import "fmt"

// BubbleSort perfoms bubble sort in ascending mode
func BubbleSort(arr []int) {
	size := len(arr)
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				/* Swapping */
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func main() {
	data := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	BubbleSort(data)
	fmt.Println(data)
}
