package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	count := make(map[int]int)
	result := []int{}

	for _, num := range nums1 {
		count[num]++
	}

	for _, num := range nums2 {
		if count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
	}

	return result
}

func intersect1(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	result := []int{}

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			// Элементы равны — пересечение найдено
			result = append(result, nums1[i])
			i++
			j++
		}
	}

	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	intersection := intersect(nums1, nums2)
	fmt.Println("Intersection:", intersection)
}
