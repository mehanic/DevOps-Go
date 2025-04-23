package main

import (
	"fmt"
	"math"
)

// Функция для поиска минимума и максимума в массиве
func FindMinMax(numbers []int) (int, int) {
	// Проверка: если массив пустой, возвращаем ошибочные значения
	if len(numbers) == 0 {
		fmt.Println("Массив пустой")
		return 0, 0
	}

	// Инициализируем min и max начальными значениями
	min := math.MaxInt // максимальное возможное число в int
	max := math.MinInt // минимальное возможное число в int

	// Проходим по массиву и обновляем min и max
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Возвращаем найденные min и max
	return min, max
}

func main() {
	// Пример использования
	arr := []int{5, 2, 9, 1, 7, -3, 8}
	min, max := FindMinMax(arr)
	fmt.Println("Минимальное число:", min)
	fmt.Println("Максимальное число:", max)
}
