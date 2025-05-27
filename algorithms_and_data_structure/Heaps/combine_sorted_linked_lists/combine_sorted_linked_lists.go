package main

import (
	"container/heap"
	"fmt"
)

// Определение ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Реализация кучи (min-heap) для ListNode
type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Функция объединения отсортированных связных списков
func combineSortedLinkedLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	// Помещаем в кучу головы всех списков
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for h.Len() > 0 {
		smallest := heap.Pop(h).(*ListNode)
		curr.Next = smallest
		curr = curr.Next
		if smallest.Next != nil {
			heap.Push(h, smallest.Next)
		}
	}

	return dummy.Next
}

// Пример использования
func main() {
	// Создаем несколько отсортированных списков
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	l3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

	merged := combineSortedLinkedLists([]*ListNode{l1, l2, l3})
	for merged != nil {
		fmt.Print(merged.Val, " ")
		merged = merged.Next
	}
}
