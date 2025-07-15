package main

import (
	"fmt"
)

// mergeSort recursively divides and sorts the list
func mergeSort(arr []int, start int, end int) []int {
	if start == end {
		return []int{arr[start]}
	}
	mid := (start + end) / 2
	left := mergeSort(arr, start, mid)
	right := mergeSort(arr, mid+1, end)
	return merge(left, right)
}

// merge combines two sorted slices into one sorted slice
func merge(p1, p2 []int) []int {
	i, j := 0, 0
	sorted := []int{}

	for i < len(p1) && j < len(p2) {
		if p1[i] <= p2[j] {
			sorted = append(sorted, p1[i])
			i++
		} else {
			sorted = append(sorted, p2[j])
			j++
		}
	}
	for i < len(p1) {
		sorted = append(sorted, p1[i])
		i++
	}
	for j < len(p2) {
		sorted = append(sorted, p2[j])
		j++
	}
	return sorted
}

func main() {
	l := []int{4, 5, 9, 8, 7, 6, 2, 3}
	sorted := mergeSort(l, 0, len(l)-1)
	fmt.Println(sorted)
}
