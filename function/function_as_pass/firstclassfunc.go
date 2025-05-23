package main

import "fmt"

// Высшего порядка функция, принимает операцию и аргументы
func operate(op func(a, b int64) int64, a, b int64) {
	fmt.Println(op(a, b))
}

func main() {
	// Разные реализации операций
	add := func(a, b int64) int64 {
		return a + b
	}

	sub := func(a, b int64) int64 {
		return a - b
	}

	mul := func(a, b int64) int64 {
		return a * b
	}

	div := func(a, b int64) int64 {
		if b == 0 {
			fmt.Println("деление на ноль")
			return 0
		}
		return a / b
	}

	// Используем одну функцию operate для разных операций
	operate(add, 10, 5) // 15
	operate(sub, 10, 5) // 5
	operate(mul, 10, 5) // 50
	operate(div, 10, 5) // 2
}

