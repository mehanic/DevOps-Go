package main

import (
	"fmt"
)

func longestSubstringWithUniqueChars(s string) int {
	maxLen := 0
	seen := make(map[rune]bool)
	left := 0
	runes := []rune(s) // Обработка Unicode символов

	for right := 0; right < len(runes); right++ {
		// Сжимаем окно, пока не удалим дубликат
		for seen[runes[right]] {
			delete(seen, runes[left])
			left++
		}
		// Обновляем максимум
		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
		seen[runes[right]] = true
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println("Longest substring with unique characters:", longestSubstringWithUniqueChars(s)) // Output: 3
}
