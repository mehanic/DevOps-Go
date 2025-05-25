package main

import "fmt"

// @leet start
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

// @leet end
func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	intersection := intersect(nums1, nums2)
	fmt.Println("Intersection:", intersection)
}
