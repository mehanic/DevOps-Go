package main

import "fmt"

func BubbleSort(input []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < len(input)-1-counter; i++ {
			if input[i] > input [i + 1] {
				input[i], input[i+1] = input[i+1], input[i]
				isSorted = false
			}
		} 
		counter++
	}
	return input
}


func main() {
    // Example 1: Unsorted array
    arr1 := []int{5, 3, 8, 4, 2}
    sortedArr1 := BubbleSort(arr1)
    fmt.Println("Sorted Array 1:", sortedArr1)
	}
