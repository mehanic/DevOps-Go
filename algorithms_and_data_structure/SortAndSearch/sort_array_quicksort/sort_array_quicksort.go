package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
	return nums
}

func quicksort(nums []int, left, right int) {
	// Базовый случай: если подмассив содержит 0 или 1 элемент, он уже отсортирован
	if left >= right {
		return
	}
	// Разделить массив и получить индекс опорного элемента
	pivotIndex := partition(nums, left, right)
	// Рекурсивно отсортировать левую и правую части
	quicksort(nums, left, pivotIndex-1)
	quicksort(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	lo := left
	// Переместить все элементы меньше опорного влево
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[lo] = nums[lo], nums[i]
			lo++
		}
	}
	// Поместить опорный элемент на правильную позицию
	nums[lo], nums[right] = nums[right], nums[lo]
	return lo
}

func main() {
	nums := []int{5, 2, 3, 1, 4}
	sorted := sortArray(nums)
	fmt.Println(sorted) // [1 2 3 4 5]
}
