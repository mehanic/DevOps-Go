package main

import (
	"fmt"
)

func findAllSubsets(nums []int) [][]int {
	var res [][]int
	var currSubset []int
	backtrack(0, currSubset, nums, &res)
	return res
}

func backtrack(i int, currSubset []int, nums []int, res *[][]int) {
	// Базовый случай: если все элементы рассмотрены — добавляем копию текущего подмножества
	if i == len(nums) {
		tmp := make([]int, len(currSubset))
		copy(tmp, currSubset)
		*res = append(*res, tmp)
		return
	}

	// Включить текущий элемент
	currSubset = append(currSubset, nums[i])
	backtrack(i+1, currSubset, nums, res)

	// Исключить текущий элемент (backtrack)
	currSubset = currSubset[:len(currSubset)-1]
	backtrack(i+1, currSubset, nums, res)
}

func main() {
	nums := []int{1, 2, 3}
	subsets := findAllSubsets(nums)
	fmt.Println(subsets)
}
