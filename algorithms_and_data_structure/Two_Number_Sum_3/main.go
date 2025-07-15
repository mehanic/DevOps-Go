package main

import "fmt"

func TwoSum(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == target {
				return []int{i, j} // Вернуть индексы, а не сами числа!
			}
		}
	}
	return []int{} // Если не найдено
}

func main() {
	array := []int{3, 4, 5, 6}
	target := 7
	results := TwoSum(array, target)
	fmt.Println(results) // Ожидаемый вывод: [0,1]

	main1()
	main3()
}

func main1() {
	array := []int{4, 5, 6}
	target := 10
	results := TwoSum(array, target)
	fmt.Println(results) // Ожидаемый вывод: [0,2]
}

func main3() {
	array := []int{5, 5}
	target := 10
	results := TwoSum(array, target)
	fmt.Println(results) // Ожидаемый вывод: [0,1]
}
