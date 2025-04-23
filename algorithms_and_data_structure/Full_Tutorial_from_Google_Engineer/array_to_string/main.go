package main

import (
	"fmt"
	"strings"
)

// Динамический массив
type Array[T any] struct {
	arr []T
	len int
}

// Метод String() для вывода массива в строку
func (a *Array[T]) String() string {
	if a.len == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[") // Добавляем открывающую скобку

	// Перебираем все элементы массива
	for i, val := range a.arr {
		if i > 0 {
			sb.WriteString(", ") // Добавляем запятую перед каждым элементом (кроме первого)
		}
		sb.WriteString(fmt.Sprintf("%v", val)) // Преобразуем элемент в строку и добавляем
	}

	sb.WriteString("]") // Закрываем массив
	return sb.String()
}

func main() {
	// Создаем массив и заполняем его числами
	arr := &Array[int]{arr: []int{10, 20, 30, 40, 50}, len: 5}

	// Выводим массив как строку
	fmt.Println(arr.String()) // Ожидаемый вывод: [10, 20, 30, 40, 50]
}
