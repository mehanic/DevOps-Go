package main

import (
	"errors"
	"fmt"
)

// Определяем структуру для динамического массива
type Array[T comparable] struct {
	arr      []T
	len      int
	capacity int
}

// Конструктор для создания нового массива с заданной емкостью
func NewArray[T comparable](capacity int) *Array[T] {
	return &Array[T]{
		arr:      make([]T, 0, capacity),
		len:      0,
		capacity: capacity,
	}
}

// Метод для добавления элемента в массив
func (a *Array[T]) Add(elem T) {
	if a.len >= a.capacity {
		a.resize() // Увеличиваем размер массива, если нужно
	}
	a.arr = append(a.arr, elem)
	a.len++
}

// Метод для увеличения емкости массива
func (a *Array[T]) resize() {
	newCapacity := a.capacity * 2
	newArr := make([]T, a.len, newCapacity)
	copy(newArr, a.arr)
	a.arr = newArr
	a.capacity = newCapacity
}

// Метод для удаления элемента по индексу
func (a *Array[T]) RemoveAt(rmIndex int) (T, error) {
	if rmIndex < 0 || rmIndex >= a.len {
		var zeroValue T
		return zeroValue, errors.New("Index out of bounds")
	}

	// Сохраняем удаляемый элемент
	data := a.arr[rmIndex]

	// Перемещаем все элементы после rmIndex на одну позицию влево
	a.arr = append(a.arr[:rmIndex], a.arr[rmIndex+1:]...)

	// Обновляем длину
	a.len--

	// Возвращаем удаленный элемент
	return data, nil
}

// Метод для удаления элемента по значению
func (a *Array[T]) Remove(obj T) bool {
	for i := 0; i < a.len; i++ {
		if a.arr[i] == obj { // Если найден нужный элемент
			a.RemoveAt(i) // Удаляем его по индексу
			return true   // Возвращаем true, так как удаление произошло
		}
	}
	return false // Если элемент не найден, возвращаем false
}

func main() {
	// Создаем новый динамический массив с начальной емкостью 8
	arr := NewArray[int](10)

	// Добавляем элементы в массив
	arr.Add(10)
	arr.Add(20)
	arr.Add(30)
	arr.Add(40)
	arr.Add(50)

	// Печатаем текущий массив
	fmt.Println("Array before removal:", arr.arr)

	// Удаляем элемент 30
	removed := arr.Remove(30)
	fmt.Println("Was 30 removed?", removed) // Выведет: true

	// Удаляем элемент 100 (его нет)
	removed = arr.Remove(100)
	fmt.Println("Was 100 removed?", removed) // Выведет: false

	// Печатаем массив после удаления
	fmt.Println("Array after removal:", arr.arr)
	fmt.Println("Size after removal:", arr.len)
}
