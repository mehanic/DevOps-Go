package main

import (
	"fmt"
)

func SieveOfEratosthenes(n int) []int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	primes := []int{}

	for i := 2; i <= n; i++ {
		if isPrime[i-1] {
			primes = append(primes, i)
			for j := i * i; j <= n; j += i {
				isPrime[j-1] = false
			}
		}
	}

	return primes
}

func main() {
	n := 50
	fmt.Println(SieveOfEratosthenes(n))
}
