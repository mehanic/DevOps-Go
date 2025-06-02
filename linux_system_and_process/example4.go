package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum =+ num
	}
	return sum
}

func main() {
	// Пример использования
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	fmt.Println("Сумма чисел:", result)
}
