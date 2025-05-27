package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WeightedRandomSelection struct {
	prefixSums []int
}

func NewWeightedRandomSelection(weights []int) *WeightedRandomSelection {
	prefixSums := make([]int, len(weights))
	prefixSums[0] = weights[0]
	for i := 1; i < len(weights); i++ {
		prefixSums[i] = prefixSums[i-1] + weights[i]
	}
	return &WeightedRandomSelection{prefixSums: prefixSums}
}

func (wrs *WeightedRandomSelection) Select() int {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(wrs.prefixSums[len(wrs.prefixSums)-1]) + 1

	left, right := 0, len(wrs.prefixSums)-1
	for left < right {
		mid := (left + right) / 2
		if wrs.prefixSums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	weights := []int{1, 3, 2}
	wrs := NewWeightedRandomSelection(weights)

	// Пример: запустим выбор 10 раз
	for i := 0; i < 10; i++ {
		fmt.Println(wrs.Select())
	}
}
