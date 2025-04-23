package main

import (
	"fmt"
	"math/bits"
)

//1. Bit Manipulation - I
// countBits calculates the number of 1's in the binary representation of each number from 0 to n
func countBits(n int) []int {
    res := make([]int, n+1)
    for num := 0; num <= n; num++ {
        one := 0
        for i := 0; i < 32; i++ {
            if num & (1 << i) != 0 {
                one++
            }
        }
        res[num] = one
    }
    return res
}

func main() {
    // Example 1: n = 4
    result1 := countBits(4)
    fmt.Println("Example 1 (n=4):", result1) // Expected: [0, 1, 1, 2, 1]

    // Example 2: n = 7
    result2 := countBits(7)
    fmt.Println("Example 2 (n=7):", result2) // Expected: [0, 1, 1, 2, 1, 2, 2, 3]
}


//2. Bit Manipulation - II

func countBits1(n int) []int {
    res := make([]int, n+1)
    for i := 1; i <= n; i++ {
        num := i
        for num != 0 {
            res[i]++
            num &= (num - 1)
        }
    }
    return res
}

//3. In-Built Function

func countBits3(n int) []int {
    res := make([]int, n+1)
    for i := 0; i <= n; i++ {
        res[i] = bits.OnesCount(uint(i))
    }
    return res
}

//4. Bit Manipulation (DP)

func countBits4(n int) []int {
    dp := make([]int, n+1)
    offset := 1

    for i := 1; i <= n; i++ {
        if offset*2 == i {
            offset = i
        }
        dp[i] = 1 + dp[i - offset]
    }
    return dp
}

//5. Bit Manipulation (Optimal)

func countBits5(n int) []int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = dp[i >> 1] + (i&1);
    }
    return dp
}