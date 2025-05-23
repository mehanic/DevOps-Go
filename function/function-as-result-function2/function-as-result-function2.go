package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	if y == 0 {
		fmt.Println("Ошибка: деление на ноль!")
		return 0 //
	}
	return x / y
}

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else if n == 3 {
		return divide
	} else {
		return multiply
	}
}

func main() {
	// Сложение
	f := selectFn(1)
	fmt.Println(f(3, 4)) // 7 (сложение)

	// Вычитание
	f = selectFn(2)
	fmt.Println(f(7, 4)) // 3 (вычитание)

	// Деление
	f = selectFn(3)
	fmt.Println(f(12, 4)) // 3 (деление)

	// Умножение
	f = selectFn(4)
	fmt.Println(f(3, 4)) // 12 (умножение
}
