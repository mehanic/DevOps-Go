package main

import "fmt"

func knapsackOptimized(cap int, weights []int, values []int) int {
	n := len(values)
	prevRow := make([]int, cap+1)

	for i := n - 1; i >= 0; i-- {
		currRow := make([]int, cap+1)
		for c := 1; c <= cap; c++ {
			if weights[i] <= c {
				include := values[i] + prevRow[c-weights[i]]
				exclude := prevRow[c]
				currRow[c] = max(include, exclude)
			} else {
				currRow[c] = prevRow[c]
			}
		}
		prevRow = currRow
	}
	return prevRow[cap]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7
	fmt.Println(knapsackOptimized(capacity, weights, values)) // Output: 9
}
