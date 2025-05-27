package main

import (
	"fmt"
)

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumCollinearPoints(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		res = max(res, maxPointsFromFocalPoint(i, points))
	}
	return res
}

func maxPointsFromFocalPoint(focalIdx int, points [][]int) int {
	slopesMap := make(map[[2]int]int)
	maxPoints := 0

	for j := 0; j < len(points); j++ {
		if j != focalIdx {
			currSlope := getSlope(points[focalIdx], points[j])
			slopesMap[currSlope]++
			maxPoints = max(maxPoints, slopesMap[currSlope])
		}
	}
	// Add 1 to include the focal point itself
	return maxPoints + 1
}

func getSlope(p1, p2 []int) [2]int {
	rise := p2[1] - p1[1]
	run := p2[0] - p1[0]

	if run == 0 {
		// Vertical line slope representation
		return [2]int{1, 0}
	}
	gcdVal := gcd(abs(rise), abs(run))

	// Reduce slope
	rise /= gcdVal
	run /= gcdVal

	// Ensure consistent sign for slope representation (run positive)
	if run < 0 {
		rise = -rise
		run = -run
	}
	return [2]int{rise, run}
}

// gcd computes the Greatest Common Divisor of two integers using Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	points := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{1, 2},
		{2, 3},
	}
	fmt.Println(maximumCollinearPoints(points)) // Output should be 4
}
