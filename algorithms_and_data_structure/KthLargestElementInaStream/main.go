package main

import (
	"container/heap"
	"fmt"
)

// MinHeap is a type of heap that we will use for our KthLargest class.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// KthLargest is the main struct that handles the functionality.
type KthLargest struct {
	k      int
	minHeap MinHeap
}

// Constructor for KthLargest.
func Constructor(k int, nums []int) KthLargest {
	// Create a min-heap with the initial numbers.
	kl := KthLargest{k: k}
	heap.Init(&kl.minHeap)

	// Add all elements to the heap, maintaining the size of the heap to be at most k.
	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

// Add method to add a value to the stream and return the kth largest element.
func (kl *KthLargest) Add(val int) int {
	// Add the new value to the heap.
	heap.Push(&kl.minHeap, val)

	// If the heap has more than k elements, remove the smallest element (root of min-heap).
	if len(kl.minHeap) > kl.k {
		heap.Pop(&kl.minHeap)
	}

	// The root of the heap is the kth largest element.
	return kl.minHeap[0]
}

// Example usage of KthLargest class.
func main() {
	k := 3
	nums := []int{1, 2, 3, 3}
	kthLargest := Constructor(k, nums)
	fmt.Println(kthLargest.Add(3))  // returns 3
	fmt.Println(kthLargest.Add(5))  // returns 3
	fmt.Println(kthLargest.Add(6))  // returns 3
	fmt.Println(kthLargest.Add(7))  // returns 5
	fmt.Println(kthLargest.Add(8))  // returns 6
}

