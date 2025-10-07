package main

import (
	"fmt"
)

func findAllPermutations(nums []int) [][]int {
	var res [][]int
	var candidate []int
	used := make(map[int]bool)
	backtrack(nums, candidate, used, &res)
	return res
}

func backtrack(nums []int, candidate []int, used map[int]bool, res *[][]int) {
	// Если permutation готова, добавляем копию
	if len(candidate) == len(nums) {
		tmp := make([]int, len(candidate))
		copy(tmp, candidate)
		*res = append(*res, tmp)
		return
	}

	for _, num := range nums {
		if !used[num] {
			candidate = append(candidate, num)
			used[num] = true

			backtrack(nums, candidate, used, res)

			// Backtrack
			candidate = candidate[:len(candidate)-1]
			delete(used, num)
		}
	}
}

func main() {
	nums := []int{1, 2, 3}
	perms := findAllPermutations(nums)
	fmt.Println(perms)
}
