package main

import (
	"fmt"
)

// Finds the maximum value and index in a slice
func findMaximum(arr []int) (int, int) {
	max := arr[0]
	maxIdx := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			maxIdx = i
		}
	}
	return max, maxIdx
}

// Maintains the max-heap property
func maxHeapify(heap []int, n int, rootIdx int) []int {
	leftIdx := rootIdx*2 + 1
	rightIdx := rootIdx*2 + 2

	var subSlice []int
	if rightIdx < n {
		subSlice = []int{heap[rootIdx], heap[leftIdx], heap[rightIdx]}
	} else if leftIdx < n {
		subSlice = []int{heap[rootIdx], heap[leftIdx]}
	} else {
		return heap // No children
	}

	maxVal, maxRelIdx := findMaximum(subSlice)
	var maxIdx int

	if maxVal != heap[rootIdx] {
		if maxRelIdx == 1 {
			maxIdx = leftIdx
		} else if maxRelIdx == 2 {
			maxIdx = rightIdx
		}

		heap[rootIdx], heap[maxIdx] = heap[maxIdx], heap[rootIdx]

		if maxIdx < n/2 {
			heap = maxHeapify(heap, n, maxIdx)
		}
	}

	return heap
}

// Builds the max-heap tree from an array
func buildHeapTree(heap []int, n int) []int {
	startIdx := n/2 - 1
	for startIdx >= 0 {
		heap = maxHeapify(heap, n, startIdx)
		startIdx--
	}
	return heap
}

// Main function: Heap Sort
func main() {
	input := []int{1, 12, 9, 5, 6, 10}
	heapifiedList := make([]int, len(input))
	copy(heapifiedList, input)

	var sortedList []int

	for len(heapifiedList) > 0 {
		heapifiedList = buildHeapTree(heapifiedList, len(heapifiedList))
		// Swap root and last element
		heapifiedList[0], heapifiedList[len(heapifiedList)-1] = heapifiedList[len(heapifiedList)-1], heapifiedList[0]
		// Pop max element and insert at front of sorted list
		sortedList = append([]int{heapifiedList[len(heapifiedList)-1]}, sortedList...)
		heapifiedList = heapifiedList[:len(heapifiedList)-1]
	}

	fmt.Println("Answer:", sortedList)
}
