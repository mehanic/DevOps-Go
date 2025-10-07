package main

import (
	"fmt"
)


func maxProfit(prices []int) int {
	profit := 0
	for elem := 1; elem < len(prices); elem++ {
		if prices[elem] > prices[elem-1] {
			profit += prices[elem] - prices[elem-1]
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max income:", maxProfit(prices)) // output 7

	prices1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Max income:", maxProfit(prices1)) // output 4

	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println("Max income:", maxProfit(prices2)) // output 0
}

