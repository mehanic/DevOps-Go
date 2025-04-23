package main

import (
	"fmt"
	"errors"
)

// Определяем структуру для нашего массива
type Array[T any] struct {
	arr      []T
	len      int
	capacity int
}

// Конструктор с дефолтной емкостью 16
func NewArray[T any](capacity int) (*Array[T], error) {
	if capacity < 0 {
		return nil, errors.New("Illegal Capacity: " + fmt.Sprint(capacity))
	}
	return &Array[T]{
		arr:      make([]T, capacity),
		len:      0,
		capacity: capacity,
	}, nil
}

// Метод для получения размера массива
func (a *Array[T]) Size() int {
	return a.len
}

// Метод для проверки, пустой ли массив
func (a *Array[T]) IsEmpty() bool {
	return a.Size() == 0
}

// Метод для получения элемента по индексу
func (a *Array[T]) Get(index int) (T, error) {
	if index < 0 || index >= a.len {
		var zeroValue T
		return zeroValue, errors.New("Index out of bounds")
	}
	return a.arr[index], nil
}

// Метод для установки элемента по индексу
func (a *Array[T]) Set(index int, elem T) error {
	if index < 0 || index >= a.len {
		return errors.New("Index out of bounds")
	}
	a.arr[index] = elem
	return nil
}

// Метод для добавления элемента в конец массива
func (a *Array[T]) Add(elem T) {
	if a.len >= a.capacity {
		a.resize()
	}
	a.arr[a.len] = elem
	a.len++
}

// Метод для изменения емкости массива (если он заполнен)
func (a *Array[T]) resize() {
	newCapacity := a.capacity * 2
	newArr := make([]T, newCapacity)
	copy(newArr, a.arr)
	a.arr = newArr
	a.capacity = newCapacity
}

func main() {
	// Создаем массив с емкостью 4 для типа int
	arr, err := NewArray[int](4)  // Здесь указываем тип для NewArray
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Добавляем элементы
	arr.Add(10)
	arr.Add(20)
	arr.Add(30)

	// Получаем размер массива
	fmt.Println("Size:", arr.Size()) // Выведет: Size: 3

	// Проверяем, пуст ли массив
	fmt.Println("Is empty?", arr.IsEmpty()) // Выведет: Is empty? false

	// Получаем элемент по индексу
	val, err := arr.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index 1:", val) // Выведет: Element at index 1: 20
	}

	// Устанавливаем новый элемент по индексу
	err = arr.Set(1, 99)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Проверяем изменение
	val, _ = arr.Get(1)
	fmt.Println("Updated element at index 1:", val) // Выведет: Updated element at index 1: 99
}