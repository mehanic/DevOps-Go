package main

import (
	"container/heap"
	"fmt"
)

// MaxHeap для отрицательных чисел, чтобы реализовать max-heap через container/heap
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // обратный порядок
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinHeap для стандартного порядка
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

type MedianOfAnIntegerStream struct {
	leftHalf  *MaxHeap // max-heap
	rightHalf *MinHeap // min-heap
}

func NewMedianOfAnIntegerStream() *MedianOfAnIntegerStream {
	left := &MaxHeap{}
	right := &MinHeap{}
	heap.Init(left)
	heap.Init(right)
	return &MedianOfAnIntegerStream{
		leftHalf:  left,
		rightHalf: right,
	}
}

func (m *MedianOfAnIntegerStream) Add(num int) {
	if m.leftHalf.Len() == 0 || num <= (*m.leftHalf)[0] {
		heap.Push(m.leftHalf, num)
		if m.leftHalf.Len() > m.rightHalf.Len()+1 {
			val := heap.Pop(m.leftHalf).(int)
			heap.Push(m.rightHalf, val)
		}
	} else {
		heap.Push(m.rightHalf, num)
		if m.leftHalf.Len() < m.rightHalf.Len() {
			val := heap.Pop(m.rightHalf).(int)
			heap.Push(m.leftHalf, val)
		}
	}
}

func (m *MedianOfAnIntegerStream) GetMedian() float64 {
	if m.leftHalf.Len() == m.rightHalf.Len() {
		return (float64((*m.leftHalf)[0]) + float64((*m.rightHalf)[0])) / 2.0
	}
	return float64((*m.leftHalf)[0])
}

func main() {
	stream := NewMedianOfAnIntegerStream()
	stream.Add(1)
	fmt.Println(stream.GetMedian()) // 1
	stream.Add(5)
	fmt.Println(stream.GetMedian()) // 3
	stream.Add(2)
	fmt.Println(stream.GetMedian()) // 2
	stream.Add(10)
	fmt.Println(stream.GetMedian()) // 3.5
}
