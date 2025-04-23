package main

import (
	"errors"
	"fmt"
)

// Определяем структуру для динамического массива
type Array[T any] struct {
	arr      []T
	len      int
	capacity int
}

// Конструктор для создания нового массива с заданной емкостью
func NewArray[T any](capacity int) *Array[T] {
	return &Array[T]{
		arr:      make([]T, 0, capacity), // Массив с начальной емкостью
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

func main() {
	// Создаем новый динамический массив с начальной емкостью 8
	arr := NewArray[int](8)

	// Добавляем элементы в массив
	arr.Add(10)
	arr.Add(20)
	arr.Add(30)
	arr.Add(40)
	arr.Add(50)

	// Печатаем текущий массив
	fmt.Println("Array before removal:", arr.arr)

	// Удаляем элемент с индексом 2 (элемент 30)
	val, err := arr.RemoveAt(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Removed element:", val) // Удаляем элемент 30
	}

	// Печатаем массив после удаления
	fmt.Println("Array after removal:", arr.arr)
	fmt.Println("Size after removal:", arr.len)
}
