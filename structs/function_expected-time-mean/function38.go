package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Function to generate raw data
func rawData(n int, limit int, arrivalFunc func(int) []int) map[int]int {
	data := arrivalFunc(n)
	waitTimes := make(map[int]int)
	for _, item := range data {
		waitTimes[item]++
	}
	return waitTimes
}

// Function to simulate arrival data (stub implementation)
func arrival1(n int) []int {
	// This is a stub. Replace with actual implementation.
	return make([]int, n)
}

// LazyCounterStatistics struct holds raw data and calculates statistics
type LazyCounterStatistics struct {
	rawCounter map[int]int
}

// Sum calculates the sum of values weighted by their frequencies
func (lcs *LazyCounterStatistics) Sum() float64 {
	sum := 0.0
	for value, freq := range lcs.rawCounter {
		sum += float64(value * freq)
	}
	return sum
}

// Count returns the total count of all frequencies
func (lcs *LazyCounterStatistics) Count() int {
	count := 0
	for _, freq := range lcs.rawCounter {
		count += freq
	}
	return count
}

// Sum2 calculates the sum of squared values weighted by their frequencies
func (lcs *LazyCounterStatistics) Sum2() float64 {
	sum := 0.0
	for value, freq := range lcs.rawCounter {
		sum += float64(value * value * freq)
	}
	return sum
}

// Mean calculates the mean of the raw data
func (lcs *LazyCounterStatistics) Mean() float64 {
	count := float64(lcs.Count())
	if count == 0 {
		return 0
	}
	return lcs.Sum() / count
}

// Variance calculates the variance of the raw data
func (lcs *LazyCounterStatistics) Variance() float64 {
	n := float64(lcs.Count())
	if n <= 1 {
		return 0
	}
	return (lcs.Sum2() - (lcs.Sum()*lcs.Sum()/n)) / (n - 1)
}

// StdDev calculates the standard deviation of the raw data
func (lcs *LazyCounterStatistics) StdDev() float64 {
	return math.Sqrt(lcs.Variance())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data := rawData(8, 1000, arrival1)

	// Print expected time and statistics
	fmt.Printf("Expected Time: %.2f\n", float64(expected(8))) // You need to implement expected() function
	fmt.Printf("Expected Mean: %.2f\n", mean(data))
	fmt.Printf("Expected StdDev: %.2f\n", stddev(data))

	stats := LazyCounterStatistics{rawCounter: data}
	fmt.Printf("Mean: %.2f\n", stats.Mean())
	fmt.Printf("Standard Deviation: %.3f\n", stats.StdDev())
}

// Stub functions for expected(), mean(), and stddev()
func expected(n int) float64 {
	// Placeholder implementation, replace with actual function
	return 0.0
}

func mean(data map[int]int) float64 {
	sum := 0.0
	count := 0
	for value, freq := range data {
		sum += float64(value * freq)
		count += freq
	}
	return sum / float64(count)
}

func stddev(data map[int]int) float64 {
	meanVal := mean(data)
	sum2 := 0.0
	count := 0
	for value, freq := range data {
		sum2 += float64(value * value * freq)
		count += freq
	}
	variance := (sum2 - meanVal*meanVal*float64(count)) / float64(count-1)
	return math.Sqrt(variance)
}

// Expected Time: 0.00
// Expected Mean: 0.00
// Expected StdDev: 0.00
// Mean: 0.00
// Standard Deviation: 0.000

