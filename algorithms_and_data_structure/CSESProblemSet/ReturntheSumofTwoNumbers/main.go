package main

import (
	"fmt"
)

// Функция, которая возвращает сумму двух чисел
func addition(a int, b int) int {
	return a + b
}

func main() {
	// Примеры
	fmt.Println(addition(3, 2))   // ➞ 5
	fmt.Println(addition(-3, -6)) // ➞ -9
	fmt.Println(addition(7, 3))   // ➞ 10
}
