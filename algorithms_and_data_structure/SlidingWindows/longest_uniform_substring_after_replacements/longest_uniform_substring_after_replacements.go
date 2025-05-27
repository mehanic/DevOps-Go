package main

import (
	"fmt"
)

func longestUniformSubstringAfterReplacements(s string, k int) int {
	freqs := make(map[byte]int)
	highestFreq := 0
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		// Увеличиваем частоту текущего символа
		freqs[s[right]]++
		if freqs[s[right]] > highestFreq {
			highestFreq = freqs[s[right]]
		}

		// Проверяем, превышает ли количество замен допустимый лимит
		windowSize := right - left + 1
		if windowSize - highestFreq > k {
			freqs[s[left]]--
			left++
		}

		// Обновляем максимальную длину
		if maxLen < right - left + 1 {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	s := "AABABBA"
	k := 1
	fmt.Println("Longest uniform substring after replacements:", longestUniformSubstringAfterReplacements(s, k)) // Output: 4
}
