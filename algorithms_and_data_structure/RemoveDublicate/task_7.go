package main

import "fmt"

func Set(x ...int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, v := range x {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

func main() {
	fmt.Println(Set([]int{1, 2, 2, 4, 5, 4, 3, 6, 7, 8, 7, 6, 9, 6, 7, 8, 9}...))
	// fmt.Println(Set([]int{1,2,2}...))
}

// []int []int{}
// [1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 4 4 4 4 4 4 4 4 4 4 4 4 5 5 5 5 5 5 5 5 5 5 5 5 4 4 4 4 4 4 4 4 4 4 4 3 3 3 3 3 3 3 3 3 3 6 6 6 6 6 6 6 7 7 7 7 7 7 8 8 8 8 8 8 7 7 7 7 7 6 6 6 6 9 9 9 6 6 6 7 7 8]

// package main

// import "fmt"

// func Set(x ...int) []int {
// 	// Используем карту для отслеживания уникальных элементов
// 	unique := make(map[int]bool)
// 	var result []int

// 	// Проходим по срезу x и добавляем в результат только уникальные элементы
// 	for _, num := range x {
// 		if !unique[num] {
// 			unique[num] = true // помечаем элемент как встреченный
// 			result = append(result, num) // добавляем в результат
// 		}
// 	}
// 	return result
// }

// func main() {
// 	// Пример использования
// 	x := []int{1, 2, 2, 3, 4, 4, 5}
// 	fmt.Println(Set(x...)) // [1, 2, 3, 4, 5]
// }
