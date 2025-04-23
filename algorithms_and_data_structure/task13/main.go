// Best Time to Buy and Sell Stock
// You are given an integer array prices where prices[i] is the price of NeetCoin on the ith day.

// You may choose a single day to buy one NeetCoin and choose a different day in the future to sell it.

// Return the maximum profit you can achieve. You may choose to not make any transactions, in which case the profit would be 0.

// Example 1:

// Input: prices = [10,1,5,6,7,1]

// Output: 6
// Explanation: Buy prices[1] and sell prices[4], profit = 7 - 1 = 6.

// Example 2:

// Input: prices = [10,8,7,5,2]

// Output: 0
// Explanation: No profitable transactions can be made, thus the max profit is 0.

// Constraints:

// 1 <= prices.length <= 100
// 0 <= prices[i] <= 100

package main

import (
	"fmt"
)

// MaxProfit function to find the maximum profit
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0] // Track the lowest price so far
	maxProfit := 0        // Track the maximum profit

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // Update min price if a lower price is found
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit // Update max profit
			}
		}
	}
	return maxProfit
}

// Main function to run test cases
func main() {
	// Example test cases
	testCases := [][]int{
		{10, 1, 5, 6, 7, 1},  // Output: 6 (Buy at 1, Sell at 7)
		{10, 8, 7, 5, 2},     // Output: 0 (No profitable transaction)
		{3, 1, 4, 8, 7, 2, 9}, // Output: 8 (Buy at 1, Sell at 9)
		{1, 2, 3, 4, 5},      // Output: 4 (Buy at 1, Sell at 5)
		{7, 6, 4, 3, 1},      // Output: 0 (No profitable transaction)
	}

	// Run test cases
	for i, prices := range testCases {
		fmt.Printf("Test Case %d: Prices: %v => Max Profit: %d\n", i+1, prices, MaxProfit(prices))
	}
}
