Sure! Below is the **Golang implementation** of the **Best Time to Buy and Sell Stock** problem. The algorithm uses a **single pass (O(n) time complexity)** to keep track of the **minimum price** seen so far and calculates the **maximum profit** efficiently.

---

### **Golang Solution**
```go
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
```

---

### **How It Works**
1. **Initialize `minPrice` to the first day's price** and `maxProfit` to `0`.
2. **Iterate through the prices**:
   - If the current price is lower than `minPrice`, update `minPrice`.
   - Otherwise, compute the profit (`price - minPrice`) and update `maxProfit` if it's higher.
3. **Return the maximum profit** found.

---

### **Example Run**
#### **Input:**
```go
prices = [10,1,5,6,7,1]
```
#### **Steps:**
- Buy at **1**, sell at **7** → **Profit = 7 - 1 = 6**.

#### **Output:**
```
Test Case 1: Prices: [10 1 5 6 7 1] => Max Profit: 6
Test Case 2: Prices: [10 8 7 5 2] => Max Profit: 0
Test Case 3: Prices: [3 1 4 8 7 2 9] => Max Profit: 8
Test Case 4: Prices: [1 2 3 4 5] => Max Profit: 4
Test Case 5: Prices: [7 6 4 3 1] => Max Profit: 0
```

---

### **Complexity Analysis**
- **Time Complexity: `O(n)`** → Only a single pass through the array.
- **Space Complexity: `O(1)`** → Only a few variables are used.

This is an **optimal solution** for this problem. Let me know if you need any modifications! 🚀