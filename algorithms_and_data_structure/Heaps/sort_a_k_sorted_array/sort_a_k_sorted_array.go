package main

import (
	"container/heap"
	"fmt"
)

// MinHeap реализует min-heap для int
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
	*h = old[:n-1]
	return x
}

func sortAKSortedArray(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	
	// Создаем мин-кучу из первых k+1 элементов
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < k+1 && i < len(nums); i++ {
		heap.Push(h, nums[i])
	}

	insertIndex := 0
	// Проходим по остальным элементам массива
	for i := k + 1; i < len(nums); i++ {
		nums[insertIndex] = heap.Pop(h).(int)
		insertIndex++
		heap.Push(h, nums[i])
	}

	// Извлекаем оставшиеся элементы из кучи
	for h.Len() > 0 {
		nums[insertIndex] = heap.Pop(h).(int)
		insertIndex++
	}

	return nums
}

func main() {
	arr := []int{3, 2, 6, 5, 4, 8}
	k := 2
	sorted := sortAKSortedArray(arr, k)
	fmt.Println(sorted) // [2 3 4 5 6 8]
}
