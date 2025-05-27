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

// MaxHeap реализует max-кучу для Pair
type MaxHeap []Pair

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	// При равенстве частот, приоритет по лексикографическому порядку (меньшая строка - выше)
	if h[i].freq == h[j].freq {
		return h[i].str < h[j].str
	}
	// Иначе, приоритет по более высокой частоте
	return h[i].freq > h[j].freq
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Функция возвращает k самых частотных строк
func kMostFrequentStringsMaxHeap(strs []string, k int) []string {
	freqs := make(map[string]int)
	for _, s := range strs {
		freqs[s]++
	}

	h := &MaxHeap{}
	heap.Init(h)
	for s, f := range freqs {
		heap.Push(h, Pair{str: s, freq: f})
	}

	result := make([]string, 0, k)
	for i := 0; i < k && h.Len() > 0; i++ {
		p := heap.Pop(h).(Pair)
		result = append(result, p.str)
	}
	return result
}

// Пример использования
func main() {
	strs := []string{"apple", "banana", "apple", "orange", "banana", "apple", "kiwi"}
	k := 2
	topK := kMostFrequentStringsMaxHeap(strs, k)
	fmt.Println(topK) // Вывод: [apple banana]
}
