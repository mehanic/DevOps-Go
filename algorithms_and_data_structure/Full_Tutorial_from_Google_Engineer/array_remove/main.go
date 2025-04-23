package main

import (
	"fmt"
	"errors"
)

type Array[T any] struct {
	arr      []T
	len      int
	capacity int
}

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

func (a *Array[T]) Size() int {
	return a.len
}

func (a *Array[T]) IsEmpty() bool {
	return a.Size() == 0
}

func (a *Array[T]) Get(index int) (T, error) {
	if index < 0 || index >= a.len {
		var zeroValue T
		return zeroValue, errors.New("Index out of bounds")
	}
	return a.arr[index], nil
}

func (a *Array[T]) Set(index int, elem T) error {
	if index < 0 || index >= a.len {
		return errors.New("Index out of bounds")
	}
	a.arr[index] = elem
	return nil
}

func (a *Array[T]) Add(elem T) {
	if a.len >= a.capacity {
		a.resize()
	}
	a.arr[a.len] = elem
	a.len++
}

func (a *Array[T]) resize() {
	newCapacity := a.capacity * 2
	newArr := make([]T, newCapacity)
	copy(newArr, a.arr)
	a.arr = newArr
	a.capacity = newCapacity
}

// Новый метод Clear для очистки массива
func (a *Array[T]) Clear() {
	// Очищаем элементы массива
	for i := 0; i < a.len; i++ {
		var zeroValue T
		a.arr[i] = zeroValue
	}
	// Сбрасываем длину массива
	a.len = 0
}

func main() {
	// Создаем массив с емкостью 4 для типа int
	arr, err := NewArray[int](8)   // Указываем емкость массива 8
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

	// Очищаем массив
	arr.Clear()

	// После очистки
	fmt.Println("Size after clear:", arr.Size()) // Выведет: Size after clear: 0
}
