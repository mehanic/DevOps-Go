package main

import (
	"fmt"
)

func sortArrayCountingSort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// Находим максимум в массиве
	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	// Создаем и заполняем массив подсчета
	counts := make([]int, maxVal+1)
	for _, num := range nums {
		counts[num]++
	}

	// Формируем отсортированный результат
	res := []int{}
	for i, count := range counts {
		for j := 0; j < count; j++ {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	nums := []int{4, 2, 2, 8, 3, 3, 1}
	sorted := sortArrayCountingSort(nums)
	fmt.Println(sorted) // [1 2 2 3 3 4 8]
}
