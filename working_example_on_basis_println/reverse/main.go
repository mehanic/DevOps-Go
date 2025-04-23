package main

import (
	"fmt"
)

func reverseSlice(slice []string) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-1-i] = slice[n-1-i], slice[i] // Меняем местами элементы
	}
}

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	reverseSlice(fruits) // Разворачиваем срез

	fmt.Println(fruits) // Вывод: [cherry banana apple]
}

