package main

import "fmt"

func neighborhoodBurglary(houses []int) int {
    n := len(houses)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return houses[0]
    }
    dp := make([]int, n)
    dp[0] = houses[0]
    dp[1] = max(houses[0], houses[1])
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], houses[i]+dp[i-2])
    }
    return dp[n-1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    houses := []int{2, 7, 9, 3, 1}
    fmt.Println(neighborhoodBurglary(houses)) // Output: 12
}
