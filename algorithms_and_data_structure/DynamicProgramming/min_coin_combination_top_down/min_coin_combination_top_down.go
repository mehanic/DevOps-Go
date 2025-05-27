package main

import (
	"fmt"
	"math"
)

func minCoinCombinationTopDown(coins []int, target int) int {
	memo := make(map[int]int)
	res := topDownDP(coins, target, memo)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func topDownDP(coins []int, target int, memo map[int]int) int {
	if target == 0 {
		return 0
	}
	if val, exists := memo[target]; exists {
		return val
	}

	minCoins := math.MaxInt32
	for _, coin := range coins {
		if coin <= target {
			// Recursive call for the remaining amount
			curr := topDownDP(coins, target-coin, memo)
			if curr != math.MaxInt32 && curr+1 < minCoins {
				minCoins = curr + 1
			}
		}
	}

	memo[target] = minCoins
	return minCoins
}

func main() {
	coins := []int{1, 2, 5}
	target := 11
	fmt.Println(minCoinCombinationTopDown(coins, target)) // Output: 3 (5 + 5 + 1)
}
