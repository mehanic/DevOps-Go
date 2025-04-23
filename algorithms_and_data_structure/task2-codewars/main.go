// Write functions:

// 1: an(n) with parameter n: returns the first n terms of the series of a(n) (not tested)

// 2: gn(n) with parameter n: returns the first n terms of the series of g(n) (not tested)

// 3: countOnes(n) with parameter n: returns the number of 1 in the series gn(n)
//     (don't forget to add a `1` at the head) # (tested)

// 4:  p(n) with parameter n: returns an array filled with the n first distinct primes in the same order they are found in the sequence gn (not tested)

// 5: maxPn(n) with parameter n: returns the biggest prime number of the above p(n) # (tested)

// 6: anOver(n) with parameter n: returns an array (n terms) of the a(i)/i for every i such g(i) != 1 (not tested but interesting result)

// 7: anOverAverage(n) with parameter n: returns as an *integer* the average of anOver(n) # (tested)
// Note:
// You can write directly functions 3:, 5: and 7:. There is no need to write functions 1:, 2:, 4: 6: except out of pure curiosity.

//////
// Time: 2344ms Passed: 3Failed: 0
// Test Results:
// Tests
// should handle basic cases CountOnes
// Test Passed
// Completed in 0.0737ms
// should handle basic cases MaxPn
// Log
// Generated g(n): [1 1 1 1 5 3 1 1 1 1 11 3 1 1 1 1 1 1 1 1 1 1 23 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 47 3 1 5 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 101 3 1 1 7 1 1 1 1 11 3 1 1 1 1 1 13 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 233 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 467 3 1 5 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 0]
// Extracted primes: [5 3 11 23 47]
// Generated g(n): [1 1 1 1 5 3 1 1 1 1 11 3 1 1 1 1 1 1 1 1 1 1 23 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 47 3 1 5 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 101 3 1 1 7 1 1 1 1 11 3 1 1 1 1 1 13 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 233 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 467 3 1 5 3 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 0]
// Extracted primes: [5 3 11 23 47 101 7]
// Test Passed
// Completed in 0.7616ms
// should handle basic cases AnOverAverage
// Test Passed
// Completed in 0.0137ms

package kata

import (
	"fmt"
	"math"
	"sort"
)

// Helper function to compute GCD
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Generate the sequence a(n)
func an(n int64) []int64 {
	a := make([]int64, n)
	a[0] = 7
	for i := int64(1); i < n; i++ {
		a[i] = a[i-1] + gcd(i+1, a[i-1])
	}
	return a
}

// Generate the sequence g(n)
func gn(n int64) []int64 {
	a := an(n)
	g := make([]int64, n) // Fix: should not be `n+1`

	g[0] = a[0] // Fix: Initialize correctly

	for i := int64(1); i < n; i++ {
		g[i] = a[i] - a[i-1]
	}

	fmt.Println("gn(n):", g) // Debug log
	return g
}

// Count occurrences of 1 in gn(n)
func CountOnes(n int64) int {
	g := gn(n)
	count := 0
	for _, v := range g {
		if v == 1 {
			count++
		}
	}
	return count
}

// Check if a number is prime
func isPrime(num int64) bool {
	if num < 2 {
		return false
	}
	for i := int64(2); i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Extract first `n` distinct primes from g(n)
func p(n int64) []int64 {
	g := gn(n * 100) // Increase factor to ensure enough primes are collected
	primes := []int64{}
	seen := make(map[int64]bool)

	fmt.Println("Generated g(n):", g) // Debugging output

	for _, v := range g {
		if v > 1 && isPrime(v) && !seen[v] {
			primes = append(primes, v)
			seen[v] = true
		}
		if int64(len(primes)) == n {
			break
		}
	}

	fmt.Println("Extracted primes:", primes) // Debugging output
	return primes
}

// Get the maximum prime from p(n)
func MaxPn(n int64) int64 {
	primes := p(n)
	if len(primes) == 0 {
		return -1
	}
	sort.Slice(primes, func(i, j int) bool { return primes[i] < primes[j] }) // Sort primes in ascending order
	return primes[len(primes)-1]                                             // Return largest prime
}

// Compute a(i) / i for g(i) != 1
func anOver(n int64) []int64 {
	a := an(n)
	g := gn(n)
	result := []int64{}

	for i := int64(1); i < n; i++ {
		if g[i] != 1 {
			result = append(result, a[i]/(i+1))
		}
	}
	return result
}

// Compute the integer average of anOver(n)
func AnOverAverage(n int64) int {
	over := anOver(n)
	if len(over) == 0 {
		fmt.Println("AnOverAverage: Empty anOver, returning 0") // Debug log
		return 0
	}

	sum := int64(0)
	for _, v := range over {
		sum += v
	}

	average := int(math.Round(float64(sum) / float64(len(over))))
	fmt.Println("AnOverAverage:", average) // Debug log
	return average
}
