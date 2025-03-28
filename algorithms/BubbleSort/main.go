package main

import "fmt"


func BubbleSort(input []int) []int {
	isSorted := false
	counter := 0
	for !isSorted {
		isSorted = true
		for i := 0; i < len(input)-1-counter; i++ {
			if input[i] > input[i+1] {
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

    // Example 2: Already sorted array
    arr2 := []int{1, 2, 3, 4, 5}
    sortedArr2 := BubbleSort(arr2)
    fmt.Println("Sorted Array 2:", sortedArr2)

    // Example 3: Array in reverse order
    arr3 := []int{9, 7, 5, 3, 1}
    sortedArr3 := BubbleSort(arr3)
    fmt.Println("Sorted Array 3:", sortedArr3)

    // Example 4: Array with repeated elements
    arr4 := []int{4, 4, 2, 1, 3, 2, 5, 1}
    sortedArr4 := BubbleSort(arr4)
    fmt.Println("Sorted Array 4:", sortedArr4)
}