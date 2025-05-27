package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func tripletSumBruteForce(nums []int) [][]int {
	n := len(nums)
	unique := make(map[string]struct{})
	result := [][]int{}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					sort.Ints(triplet) // сортировка для уникальности

					// ключ: строка вида "-1,0,1"
					key := tripletToString(triplet)
					if _, exists := unique[key]; !exists {
						unique[key] = struct{}{}
						result = append(result, triplet)
					}
				}
			}
		}
	}
	return result
}

func tripletToString(triplet []int) string {
	strs := make([]string, len(triplet))
	for i, num := range triplet {
		strs[i] = strconv.Itoa(num)
	}
	return strings.Join(strs, ",")
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(tripletSumBruteForce(nums))
}
