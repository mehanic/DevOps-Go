package main

import (
	"fmt"
)

func InsertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

func main() {
	// Несколько тестовых массивов
	testCases := [][]int{
		{5, 3, 8, 4, 2},
		{1, 2, 3, 4, 5},
		{9, 7, 5, 3, 1},
		{4, 4, 4, 4, 4},
		{-3, -1, -7, -5, -2},
		{10, 3, 3, 7, 8, 3, 2},
	}

	// Применение сортировки ко всем тестовым случаям
	for _, arr := range testCases {
		fmt.Println("Исходный массив:", arr)
		sortedArr := InsertionSort(arr)
		fmt.Println("Отсортированный массив:", sortedArr)
		fmt.Println()
	}
}



