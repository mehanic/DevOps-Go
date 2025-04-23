package main

import (
	"fmt"
)

// Определяем структуру для динамического массива
type Array[T comparable] struct {
	arr []T
	len int
}

// Метод indexOf ищет индекс элемента в массиве
func (a *Array[T]) IndexOf(obj T) int {
	for i := 0; i < a.len; i++ {
		if a.arr[i] == obj { // Если элемент найден, возвращаем индекс
			return i
		}
	}
	return -1 // Если элемент не найден, возвращаем -1
}

// Метод contains проверяет, есть ли элемент в массиве
func (a *Array[T]) Contains(obj T) bool {
	return a.IndexOf(obj) != -1 // Если indexOf не вернул -1, значит, элемент есть
}

func main() {
	// Создаем массив и заполняем его числами
	arr := &Array[int]{arr: []int{10, 20, 30, 40, 50}, len: 5}

	// Проверяем индекс элемента
	index := arr.IndexOf(30)
	fmt.Println("Index of 30:", index) // Должно вывести: Index of 30: 2

	// Проверяем, есть ли элемент в массиве
	fmt.Println("Contains 30:", arr.Contains(30)) // true
	fmt.Println("Contains 100:", arr.Contains(100)) // false
}
