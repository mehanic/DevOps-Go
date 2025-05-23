package main

import (
	"fmt"
)

// Описание типа операции (функция, принимающая два int64 и возвращающая int64)
type binaryOp func(a, b int64) int64

// Функция operate теперь возвращает результат
func operate(op binaryOp, name string, a, b int64) int64 {
	result := op(a, b)
	fmt.Printf("Операция %s: %d и %d → %d\n", name, a, b, result)
	return result
}

func main() {
	// Операции как функции
	add := func(a, b int64) int64 { return a + b }
	sub := func(a, b int64) int64 { return a - b }
	mul := func(a, b int64) int64 { return a * b }
	div := func(a, b int64) int64 {
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return 0
		}
		return a / b
	}

	// Вызов с разными операциями
	operate(add, "сложение", 10, 5)
	operate(sub, "вычитание", 10, 5)
	operate(mul, "умножение", 10, 5)
	operate(div, "деление", 10, 5)
}
