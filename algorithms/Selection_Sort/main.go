package main

import (
	"fmt"
)

// Selection Sort
func SelectionSort(array []int) []int {
	for currentIndex := 0; currentIndex < len(array)-1; currentIndex++ {
		smallestIndex := currentIndex
		for i := currentIndex + 1; i < len(array); i++ {
			if array[smallestIndex] > array[i] {
				smallestIndex = i
			}
		}
		// Меняем элементы местами
		array[currentIndex], array[smallestIndex] = array[smallestIndex], array[currentIndex]
	}
	return array
}

func main() {
	// Примеры массивов
	testCases := [][]int{
		{29, 10, 14, 37, 13},
		{3, 1, 2, 5, 4},
		{9, 7, 5, 3, 1},
		{4, 4, 4, 4, 4},
		{10, -3, 0, 5, -1},
	}

	// Применение сортировки и вывод результата
	for _, arr := range testCases {
		fmt.Println("Исходный массив:", arr)
		sortedArr := SelectionSort(arr)
		fmt.Println("Отсортированный массив:", sortedArr)
		fmt.Println()
	}
}
