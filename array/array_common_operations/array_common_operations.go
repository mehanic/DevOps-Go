package main

import (
	"fmt"
)

func modifyArray(arr [5]int, index int, value int) [5]int {
	arr[index] = value
	return arr
}

func main() {
	originalArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Length of the array is:", len(originalArray))

	copiedArray := originalArray
	fmt.Println("Copied Array:", copiedArray)

	modifiedArray := modifyArray(originalArray, 2, 10)
	fmt.Println("Modified Array:", modifiedArray)

	isEqual := originalArray == modifiedArray
	fmt.Println("Are original and modified arrays equal?", isEqual)

	isCopiedEqual := originalArray == copiedArray
	fmt.Println("Are original and copied arrays equal?", isCopiedEqual)
}
