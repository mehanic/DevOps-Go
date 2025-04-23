package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Работа с срезом ===")
	n := 5
	a := []int{}
	fmt.Println(len(a))
	fmt.Printf("Длина среза ДО добавления элементов: %d\n", len(a))

	for i := 0; i < n; i++ {
		k := 5
		a = append(a, k)
	}
	fmt.Println(len(a))
	fmt.Printf("Длина среза ПОСЛЕ добавления элементов: %d\n", len(a))

	// Опционально: показать сам срез
	strSlice := []string{}
	for _, v := range a {
		strSlice = append(strSlice, fmt.Sprintf("%d", v))
	}
	fmt.Println("Содержимое среза:", strings.Join(strSlice, ", "))
}
