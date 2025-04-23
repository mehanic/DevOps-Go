package main

import "fmt"

// Структура, которая будет хранить слайс и информацию о его длине и емкости
type DynamicArray[T any] struct {
	arr      []T
	len      int
	capacity int
}

// Конструктор для создания нового динамического массива
func NewDynamicArray[T any](initialCapacity int) *DynamicArray[T] {
	if initialCapacity <= 0 {
		initialCapacity = 1
	}
	return &DynamicArray[T]{
		arr:      make([]T, initialCapacity),
		len:      0,
		capacity: initialCapacity,
	}
}

// Метод для добавления элемента в массив
func (d *DynamicArray[T]) Add(elem T) {
	// Если места в массиве нет, увеличиваем его емкость
	if d.len >= d.capacity {
		// Увеличиваем емкость в два раза
		d.capacity *= 2

		// Создаем новый слайс с увеличенной емкостью
		newArr := make([]T, d.capacity)

		// Копируем элементы из старого слайса в новый
		copy(newArr, d.arr)

		// Переназначаем слайс arr на новый слайс с увеличенной емкостью
		d.arr = newArr
	}

	// Добавляем новый элемент в слайс
	d.arr[d.len] = elem
	d.len++
}

// Метод для получения размера слайса
func (d *DynamicArray[T]) Size() int {
	return d.len
}

// Метод для получения элемента по индексу
func (d *DynamicArray[T]) Get(index int) (T, error) {
	if index < 0 || index >= d.len {
		var zeroValue T
		return zeroValue, fmt.Errorf("index out of bounds")
	}
	return d.arr[index], nil
}

func main() {
	// Создаем новый динамический массив с начальной емкостью 
	arr := NewDynamicArray[int](8)

	// Добавляем элементы в массив
	arr.Add(10)
	arr.Add(20)
	arr.Add(30)

	// Печатаем размер массива
	fmt.Println("Size:", arr.Size()) // Выведет: Size: 3

	// Получаем элемент по индексу
	val, err := arr.Get(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index 1:", val) // Выведет: Element at index 1: 20
	}

	// Добавляем еще элементы, чтобы проверить увеличение емкости
	arr.Add(40)
	arr.Add(50)

	// Печатаем новый размер массива
	fmt.Println("Size after adding more elements:", arr.Size()) // Выведет: Size after adding more elements: 5
}
