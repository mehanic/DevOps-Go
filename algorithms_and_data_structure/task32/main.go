package main

import (
	"fmt"
	"sort"
)
//1. Sorting
// KthLargest is a struct that tracks the kth largest number in a stream.
type KthLargest struct {
	k   int      // The k-th largest number to be tracked
	arr []int    // Array storing the added numbers
}

// Constructor initializes the KthLargest struct.
func Constructor(k int, nums []int) KthLargest {
	// Sort the array initially so that we can access the k-th largest element
	sort.Ints(nums)
	return KthLargest{k: k, arr: nums}
}

// Add method adds a new number to the collection and returns the k-th largest element.
func (this *KthLargest) Add(val int) int {
	// Append the new value to the array
	this.arr = append(this.arr, val)

	// Sort the array to keep it in ascending order
	sort.Ints(this.arr)

	// Return the k-th largest element, which is at position len(arr)-k
	return this.arr[len(this.arr)-this.k]
}

func main() {
	// Example 1: Initialize with k = 3, and nums = [4, 5, 8, 2]
	kthLargest := Constructor(3, []int{4, 5, 8, 2})

	// Add 3 to the collection
	fmt.Println(kthLargest.Add(3)) // Expected output: 4

	// Add 5 to the collection
	fmt.Println(kthLargest.Add(5)) // Expected output: 5

	// Add 10 to the collection
	fmt.Println(kthLargest.Add(10)) // Expected output: 5

	// Add 9 to the collection
	fmt.Println(kthLargest.Add(9)) // Expected output: 8

	// Add 4 to the collection
	fmt.Println(kthLargest.Add(4)) // Expected output: 8
}


