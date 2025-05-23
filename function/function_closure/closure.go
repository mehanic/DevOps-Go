package main

import (
	"fmt"
)

func closure(nums []int, sum func(l []int)) error {
	sum(nums)
	return nil
}

func main() {
	// Создаём список чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Вызываем функцию closure, передавая список и анонимную функцию
	err := closure(numbers, func(l []int) {
		total := 0
		for _, v := range l {
			total += v
		}
		fmt.Println("Сумма:", total)
	})

	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
