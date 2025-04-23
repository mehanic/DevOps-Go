// Last Stone Weight
// You are given an array of integers stones where stones[i] represents the weight of the ith stone.

// We want to run a simulation on the stones as follows:

// At each step we choose the two heaviest stones, with weight x and y and smash them togethers
// If x == y, both stones are destroyed
// If x < y, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
// Continue the simulation until there is no more than one stone remaining.

// Return the weight of the last remaining stone or return 0 if none remain.

// Example 1:

// Input: stones = [2,3,6,2,4]

// Output: 1
// Explanation:
// We smash 6 and 4 and are left with a 2, so the array becomes [2,3,2,2].
// We smash 3 and 2 and are left with a 1, so the array becomes [1,2,2].
// We smash 2 and 2, so the array becomes [1].

// Example 2:

// Input: stones = [1,2]

// Output: 1
// Constraints:

// 1 <= stones.length <= 20
// 1 <= stones[i] <= 100


package main

import (
	"container/heap"
	"fmt"
)

// Create a max-heap by using a min-heap with negative values
type MaxHeap []int

// Implement heap.Interface for MaxHeap
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // This is for max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	// Create a max heap
	h := &MaxHeap{}
	heap.Init(h)

	// Push all stones into the heap (negating the values for max-heap simulation)
	for _, stone := range stones {
		heap.Push(h, -stone)
	}

	// Run the simulation until there's no more than one stone left
	for h.Len() > 1 {
		// Pop two largest stones (smallest negative values)
		stone1 := -heap.Pop(h).(int)
		stone2 := -heap.Pop(h).(int)

		if stone1 != stone2 {
			// If the stones are different, push the difference back into the heap
			heap.Push(h, -(stone1 - stone2))
		}
	}

	// If there is a stone left, return its value; otherwise, return 0
	if h.Len() > 0 {
		return -heap.Pop(h).(int)
	}
	return 0
}

func main() {
	// Test cases
	fmt.Println(lastStoneWeight([]int{2, 3, 6, 2, 4})) // Output: 1
	fmt.Println(lastStoneWeight([]int{1, 2}))         // Output: 1
	fmt.Println(lastStoneWeight([]int{10, 10, 10, 10})) // Output: 0
	fmt.Println(lastStoneWeight([]int{5, 5, 3, 3, 2}))  // Output: 2
}
