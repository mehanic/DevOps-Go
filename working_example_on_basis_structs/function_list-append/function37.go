package main

import (
	"fmt"
	"math"
)

// StatsList is a slice of float64 numbers with statistical methods.
type StatsList []float64

// Sum returns the sum of the elements in the StatsList.
func (s StatsList) Sum() float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}

// Count returns the count of elements in the StatsList.
func (s StatsList) Count() int {
	return len(s)
}

// Sum2 returns the sum of the squares of the elements in the StatsList.
func (s StatsList) Sum2() float64 {
	sum := 0.0
	for _, v := range s {
		sum += v * v
	}
	return sum
}

// Mean returns the mean (average) of the elements in the StatsList.
func (s StatsList) Mean() float64 {
	if s.Count() == 0 {
		return 0
	}
	return s.Sum() / float64(s.Count())
}

// Variance returns the variance of the elements in the StatsList.
func (s StatsList) Variance() float64 {
	n := float64(s.Count())
	if n <= 1 {
		return 0
	}
	return (s.Sum2() - (s.Sum()*s.Sum()/n)) / (n - 1)
}

// StdDev returns the standard deviation of the elements in the StatsList.
func (s StatsList) StdDev() float64 {
	return math.Sqrt(s.Variance())
}

func main() {
	subset1 := StatsList{10, 8, 13, 9, 11}
	data := StatsList{14, 6, 4, 12, 7, 5}
	data = append(data, subset1...) // Extend the data slice with subset1

	fmt.Println(data)                           // Prints the combined data
	fmt.Printf("Mean: %.1f\n", data.Mean())     // Prints the mean
	fmt.Printf("Variance: %.1f\n", data.Variance()) // Prints the variance
	fmt.Printf("StdDev: %.1f\n", data.StdDev()) // Prints the standard deviation
}

// [14 6 4 12 7 5 10 8 13 9 11]
// Mean: 9.0
// Variance: 11.0
// StdDev: 3.3

