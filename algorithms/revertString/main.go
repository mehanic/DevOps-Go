package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s) // Преобразуем строку в срез рун (учитывает Unicode)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i] // Обмен элементов
	}

	return string(runes) // Преобразуем срез рун обратно в строку
}

func main() {
	input := "hello"
	reversed := reverseString(input)
	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reversed)
}
