// Counting Bits
// Given an integer n, count the number of 1's in the binary representation of every number in the range [0, n].

// Return an array output where output[i] is the number of 1's in the binary representation of i.

// Example 1:

// Input: n = 4

// Output: [0,1,1,2,1]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100

// Constraints:

// 0 <= n <= 1000

package main

import "fmt"

func countBits(n int) []int {
    // Create an array to store the number of 1's for each number from 0 to n
    dp := make([]int, n+1)
    
    for i := 1; i <= n; i++ {
        dp[i] = dp[i/2] + (i % 2)
    }

    return dp
}

func main() {
    // Example 1: n = 4
    result := countBits(4)
    fmt.Println(result) // Output: [0, 1, 1, 2, 1]
    
    // Example 2: n = 7
    result2 := countBits(7)
    fmt.Println(result2) // Output: [0, 1, 1, 2, 1, 2, 2, 3]
}
