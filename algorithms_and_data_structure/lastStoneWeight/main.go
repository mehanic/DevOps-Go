package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		cur := stones[len(stones)-1] - stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if cur > 0 {
			stones = append(stones, cur)
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	result := lastStoneWeight(stones)
	fmt.Println("Last stone weight:", result) // Output: 1
}
