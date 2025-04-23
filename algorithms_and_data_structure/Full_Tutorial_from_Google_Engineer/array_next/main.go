package main

import "fmt"

// Динамический массив
type Array[T any] struct {
	arr []T
	len int
}

// Итератор для массива
type Iterator[T any] struct {
	array *Array[T]
	index int
}

// Метод проверки наличия следующего элемента
func (it *Iterator[T]) HasNext() bool {
	return it.index < it.array.len
}

// Метод получения следующего элемента
func (it *Iterator[T]) Next() (T, bool) {
	if !it.HasNext() {
		var zeroValue T
		return zeroValue, false // Если элементов больше нет, возвращаем false
	}
	val := it.array.arr[it.index]
	it.index++ // Переход к следующему элементу
	return val, true
}

// Метод для получения итератора
func (a *Array[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{array: a, index: 0}
}

func main() {
	// Создаем массив и добавляем в него элементы
	arr := &Array[int]{arr: []int{10, 20, 30, 40, 50}, len: 5}

	// Получаем итератор
	it := arr.Iterator()

	// Обходим массив с помощью итератора
	for it.HasNext() {
		val, _ := it.Next()
		fmt.Println("Next element:", val)
	}
}
