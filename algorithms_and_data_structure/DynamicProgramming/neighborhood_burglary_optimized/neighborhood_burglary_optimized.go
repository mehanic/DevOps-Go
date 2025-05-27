package main

import "fmt"

func neighborhoodBurglaryOptimized(houses []int) int {
    n := len(houses)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return houses[0]
    }
    prevPrevMaxProfit := houses[0]
    prevMaxProfit := max(houses[0], houses[1])
    for i := 2; i < n; i++ {
        currMaxProfit := max(prevMaxProfit, houses[i]+prevPrevMaxProfit)
        prevPrevMaxProfit = prevMaxProfit
        prevMaxProfit = currMaxProfit
    }
    return prevMaxProfit
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    houses := []int{2, 7, 9, 3, 1}
    fmt.Println(neighborhoodBurglaryOptimized(houses)) // Output: 12
}
