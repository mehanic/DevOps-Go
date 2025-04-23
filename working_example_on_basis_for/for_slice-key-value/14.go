package main

import "fmt"

func main() {
	numbers := []int{}
	n := 100
	fmt.Println("Создаём массив чётных чисел от 0 до", n)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			numbers = append(numbers, i)
			fmt.Printf("Добавлено число: %d\n", i)
		}
	}
	fmt.Println("\nРезультат массива чётных чисел:")
	fmt.Println(numbers)

	indexes := []int{2, 6, 3}
	fmt.Println("\nИндексы, по которым ищем значения (с учётом -1):", indexes)

	for _, v := range indexes {
		targetIndex := v - 1
		fmt.Printf("\nИщем значение по индексу: %d\n", targetIndex)
		for j, value := range numbers {
			if v-1 == j {
				fmt.Print(value, " ")
			}
		}
	}
	// 0 2 4 6 8 10 12 ...
	// 2 6 3
	// 2 10 4
}
