package main

import (
	"container/heap"
	"fmt"
)

// Pair хранит строку и её частоту
type Pair struct {
	str  string
	freq int
}

// MinHeap реализует мин-кучу для Pair
type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	// При равенстве частот, приоритет по обратному лексикографическому порядку (строка "больше" идет раньше)
	if h[i].freq == h[j].freq {
		return h[i].str > h[j].str
	}
	// Иначе приоритет по меньшей частоте
	return h[i].freq < h[j].freq
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Возвращает k наиболее частотных строк, используя мин-кучу
func kMostFrequentStringsMinHeap(strs []string, k int) []string {
	freqs := make(map[string]int)
	for _, s := range strs {
		freqs[s]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for s, f := range freqs {
		heap.Push(h, Pair{str: s, freq: f})
		if h.Len() > k {
			heap.Pop(h) // Удаляем элемент с минимальной частотой/приоритетом
		}
	}

	res := make([]string, 0, k)
	for h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		res = append(res, p.str)
	}

	// Мин-куча возвращает строки в порядке от менее частых к более частым,
	// поэтому переворачиваем срез, чтобы получить правильный порядок.
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func main() {
	strs := []string{"apple", "banana", "apple", "orange", "banana", "apple", "kiwi"}
	k := 2
	topK := kMostFrequentStringsMinHeap(strs, k)
	fmt.Println(topK) // Вывод: [apple banana]
}
