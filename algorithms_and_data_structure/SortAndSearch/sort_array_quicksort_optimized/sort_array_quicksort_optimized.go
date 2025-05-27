package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sortArray(nums []int) []int {
	quicksortOptimized(nums, 0, len(nums)-1)
	return nums
}

func quicksortOptimized(nums []int, left, right int) {
	if left >= right {
		return
	}
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(right-left+1) + left
	// Перемещаем случайный опорный элемент в конец
	nums[randomIndex], nums[right] = nums[right], nums[randomIndex]

	pivotIndex := partition(nums, left, right)
	quicksortOptimized(nums, left, pivotIndex-1)
	quicksortOptimized(nums, pivotIndex+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	lo := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[lo] = nums[lo], nums[i]
			lo++
		}
	}
	nums[lo], nums[right] = nums[right], nums[lo]
	return lo
}

func main() {
	nums := []int{5, 2, 3, 1, 4}
	sorted := sortArray(nums)
	fmt.Println(sorted) // Пример: [1 2 3 4 5]
}
