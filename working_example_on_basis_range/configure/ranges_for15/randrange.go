package main

import (
	"fmt"
	"math/rand"
//	"time"
	"math/big"
//	"sort"
)

// arrival1 generates random values from 0 to n-1
func arrival1(n int) chan int {
	out := make(chan int)
	go func() {
		for {
			out <- rand.Intn(n)
		}
	}()
	return out
}

// arrival2 generates values based on a random walk with [-1, 0, +1] steps
func arrival2(n int) chan int {
	out := make(chan int)
	go func() {
		p := 0
		for {
			step := rand.Intn(3) - 1 // Choose from [-1, 0, +1]
			p += step
			if p < 0 {
				p = 0
			}
			out <- p % n
		}
	}()
	return out
}

// samples takes `limit` number of samples from a generator and returns them
func samples(limit int, generator chan int) []int {
	result := make([]int, 0, limit)
	for i := 0; i < limit; i++ {
		result = append(result, <-generator)
	}
	return result
}

// expected computes the expected value for the coupon collector's problem
func expected(n int) *big.Rat {
	sum := new(big.Rat)
	for i := 1; i <= n; i++ {
		sum.Add(sum, new(big.Rat).Inv(big.NewRat(int64(i), 1)))
	}
	result := new(big.Rat).Mul(big.NewRat(int64(n), 1), sum)
	return result
}

// couponCollector collects unique coupons from a set of `n` and returns the counts
func couponCollector(n int, data []int) []int {
	waitTimes := []int{}
	count := 0
	collection := map[int]bool{}

	for _, item := range data {
		count++
		collection[item] = true
		if len(collection) == n {
			waitTimes = append(waitTimes, count)
			count = 0
			collection = map[int]bool{}
		}
	}
	return waitTimes
}

// summary computes and prints the results for the coupon collector's problem
func summary(n, limit int, arrivalFunc func(int) chan int) {
	expectedTime := expected(n)
	data := samples(limit, arrivalFunc(n))
	waitTimes := couponCollector(n, data)
	averageTime := float64(sum(waitTimes)) / float64(len(waitTimes))

	fmt.Printf("Coupon collection, n=%d\n", n)
	fmt.Printf("Arrivals per '%T'\n", arrivalFunc)
	fmt.Printf("Expected = %s\n", expectedTime.FloatString(2))
	fmt.Printf("Actual = %.2f\n", averageTime)
}
// sum computes the sum of an integer slice
func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func main() {
	// Setting a seed for reproducibility
	rand.Seed(1)

	// Summary for arrival1
	summary(8, 1000, arrival1)

	// Summary for arrival2
	summary(8, 1000, arrival2)
}

