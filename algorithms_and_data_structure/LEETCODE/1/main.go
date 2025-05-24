// @leet start
// Remove Duplicates from Sorted Array
package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	first := 0

	for elem := 1; elem < len(nums); elem++ {
		if nums[elem] != nums[first] {
			first++
			nums[first] = nums[elem]
		}
	}

	return first + 1
}

// @leet end
func main() {
	nums := []int{1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Println("Count of unique values:", k)
	fmt.Println("Unique values:", nums[:k])
}

// ------ thicond
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	seen := make(map[int]bool)
	index := 0

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			nums[index] = num
			index++
		}
	}

	return index
}

func main1() {
	arr := []int{0, 0, 1, 1, 2, 3, 3, 4}
	n := removeDuplicates1(arr)
	fmt.Println("Unique count:", n)
	fmt.Println("Modified array:", arr[:n])
}
