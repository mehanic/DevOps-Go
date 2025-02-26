package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}
func isPositive(n int) bool {
	return n > 0
}

func sum(numbers []int, criteria func(int) bool) int {

	result := 0
	for _, val := range numbers {
		if criteria(val) {
			result += val
		}
	}
	return result
}
func main() {

	slice := []int{-2, 4, 3, -1, 7, -4, 23}

	sumOfEvens := sum(slice, isEven) // сумма четных чисел
	fmt.Println(sumOfEvens)          // -2

	sumOfPositives := sum(slice, isPositive) // сумма положительных чисел
	fmt.Println(sumOfPositives)              // 37
}

