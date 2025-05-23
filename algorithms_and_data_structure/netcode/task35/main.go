package main

import (
	"fmt"
	"math"
)

// climbStairs function calculates the number of ways to reach the nth stair
func climbStairs(n int) int {
    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n {
            if i == n {
                return 1
            }
            return 0
        }
        return dfs(i + 1) + dfs(i + 2)
    }
    return dfs(0)
}
//1. Recursion
func main() {
    // Example 1: Climbing 2 stairs
    n1 := 2
    fmt.Printf("Ways to climb %d stairs: %d\n", n1, climbStairs(n1)) // Expected output: 2

    // Example 2: Climbing 3 stairs
    n2 := 3
    fmt.Printf("Ways to climb %d stairs: %d\n", n2, climbStairs(n2)) // Expected output: 3

    // Example 3: Climbing 4 stairs
    n3 := 4
    fmt.Printf("Ways to climb %d stairs: %d\n", n3, climbStairs(n3)) // Expected output: 5

    // Example 4: Climbing 5 stairs
    n4 := 5
    fmt.Printf("Ways to climb %d stairs: %d\n", n4, climbStairs(n4)) // Expected output: 8

    // Example 5: Climbing 6 stairs
    n5 := 6
    fmt.Printf("Ways to climb %d stairs: %d\n", n5, climbStairs(n5)) // Expected output: 13
}


//2. Dynamic Programming (Top-Down)
func climbStairs1(n int) int {
    cache := make([]int, n+1)
    for i := 0; i <= n; i++ {
        cache[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n {
            if i == n {
                return 1
            }
            return 0
        }
        if cache[i] != -1 {
            return cache[i]
        }
        cache[i] = dfs(i + 1) + dfs(i + 2)
        return cache[i]
    }
    return dfs(0)
}


//3. Dynamic Programming (Bottom-Up)

func climbStairs2(n int) int {
    if n <= 2 {
        return n
    }
    dp := make([]int, n+1)
    dp[1] = 1
    dp[2] = 2
    for i := 3; i <= n; i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    return dp[n]
}

//4. Dynamic Programming (Space Optimized)

func climbStairs3(n int) int {
    one := 1
    two := 1

    for i := 0; i < n-1; i++ {
        temp := one
        one += two
        two = temp
    }

    return one
}


//5. Matrix Exponentiation

func climbStairs4(n int) int {
    if n == 1 {
        return 1
    }
    
    M := [][]int{{1, 1}, {1, 0}}
    result := matrixPow(M, n)
    
    return result[0][0]
}

func matrixMult(A, B [][]int) [][]int {
    return [][]int{
        {A[0][0]*B[0][0] + A[0][1]*B[1][0], 
         A[0][0]*B[0][1] + A[0][1]*B[1][1]},
        {A[1][0]*B[0][0] + A[1][1]*B[1][0], 
         A[1][0]*B[0][1] + A[1][1]*B[1][1]},
    }
}

func matrixPow(M [][]int, p int) [][]int {
    result := [][]int{{1, 0}, {0, 1}}
    base := M

    for p > 0 {
        if p%2 == 1 {
            result = matrixMult(result, base)
        }
        base = matrixMult(base, base)
        p /= 2
    }

    return result
}

//6. Math

func climbStairs5(n int) int {
    sqrt5 := math.Sqrt(5)
    phi := (1 + sqrt5) / 2
    psi := (1 - sqrt5) / 2
    n++
    return int(math.Round((math.Pow(phi, float64(n)) - 
               math.Pow(psi, float64(n))) / sqrt5))
}