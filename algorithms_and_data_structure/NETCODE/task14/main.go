package main

import (
	"fmt"
	"math"
)

// Алгоритм Brute Force (O(n^2))
func maxProfitBruteForce(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		buy := prices[i]
		for j := i + 1; j < len(prices); j++ {
			sell := prices[j]
			res = max(res, sell-buy) // Считаем прибыль
		}
	}
	return res
}

// Вспомогательная функция max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Главная функция для тестов
func main() {
	testCases := [][]int{
		{10, 1, 5, 6, 7, 1},   // Ответ: 6 (покупаем по 1, продаем по 7)
		{10, 8, 7, 5, 2},      // Ответ: 0 (цены падают, прибыли нет)
		{3, 1, 4, 8, 7, 2, 9}, // Ответ: 8 (покупаем по 1, продаем по 9)
		{1, 2, 3, 4, 5},       // Ответ: 4 (покупаем по 1, продаем по 5)
		{7, 6, 4, 3, 1},       // Ответ: 0 (цены падают, прибыли нет)
	}

	for i, prices := range testCases {
		fmt.Printf("Тест %d: %v -> Максимальная прибыль: %d\n", i+1, prices, maxProfitBruteForce(prices))
	}
	fmt.Println("----------------------")
	main1()
	fmt.Println("----------------------")
	main2()
}

//2. Two Pointers

// Оптимизированный алгоритм (O(n)) - Два указателя
func maxProfit1(prices []int) int {
	l, r := 0, 1 // l - покупка, r - продажа
	maxP := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l] // Вычисляем прибыль
			if profit > maxP {
				maxP = profit // Обновляем максимальную прибыль
			}
		} else {
			l = r // Обновляем минимальную цену покупки
		}
		r++ // Двигаем указатель продажи
	}
	return maxP
}

// Главная функция для тестов
func main1() {
	testCases := [][]int{
		{4, 1, 5, 6, 7, 20},    // Ожидаемый результат: 6 (покупка по 1, продажа по 7)
		{2, 8, 7, 5, 15},       // Ожидаемый результат: 0 (цены падают, нет прибыли)
		{3, 1, 4, 8, 7, 2, 15}, // Ожидаемый результат: 8 (покупка по 1, продажа по 9)
		{1, 2, 3, 4, 8},        // Ожидаемый результат: 4 (покупка по 1, продажа по 5)
		{1, 6, 4, 3, 30},       // Ожидаемый результат: 0 (цены падают, нет прибыли)
	}

	for i, prices := range testCases {
		fmt.Printf("Тест %d: %v -> Максимальная прибыль: %d\n", i+1, prices, maxProfit1(prices))
	}
}

//3. Dynamic Programming

func maxProfit2(prices []int) int {
	maxP := 0
	minBuy := math.MaxInt32

	for _, sell := range prices {
		if sell-minBuy > maxP {
			maxP = sell - minBuy
		}
		if sell < minBuy {
			minBuy = sell
		}
	}
	return maxP
}

func main2() {
	testCases := [][]int{
		{1, 1, 5, 6, 7, 100},    // Ожидаемый результат: 6 (покупка по 1, продажа по 7)
		{1, 8, 7, 5, 130},       // Ожидаемый результат: 0 (цены падают, нет прибыли)
		{45, 1, 4, 8, 7, 2, 3}, // Ожидаемый результат: 8 (покупка по 1, продажа по 9)
		{1, 2, 3, 4, 56},        // Ожидаемый результат: 4 (покупка по 1, продажа по 5)
		{12, 6, 4, 3, 1},       // Ожидаемый результат: 0 (цены падают, нет прибыли)
	}

	for i, prices := range testCases {
		fmt.Printf("Тест %d: %v -> Максимальная прибыль: %d\n", i+1, prices, maxProfit2(prices))
	}
}