package main

import (
	"fmt"
)

func candies(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	// Каждый ребёнок начинает с 1 конфетой
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// Первый проход слева направо
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// Второй проход справа налево
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i] < candies[i+1]+1 {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	// Суммируем все конфеты
	total := 0
	for _, c := range candies {
		total += c
	}

	return total
}

func main() {
	ratings := []int{1, 0, 2}
	fmt.Println("Minimum candies needed:", candies(ratings)) // Output: 5
}
