package main

import "fmt"

func hammingWeightsOfIntegers(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = countSetBits(i)
	}
	return result
}

func countSetBits(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func main() {
	n := 5
	fmt.Println(hammingWeightsOfIntegers(n)) // Output: [0 1 1 2 1 2]
}
