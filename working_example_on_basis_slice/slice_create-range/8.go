package main

import "fmt"

func main() {
	a := [][]int{}
	for i := 0; i < 4; i++ {
		arr := []int{1, 1, 1, 1}
		a = append(a, arr)
	}
	for _, v := range a {
		fmt.Println(v)
	}
}

// initialization of 2D Slice:

// The code initializes a as an empty 2D slice of integers ([][]int).
// Filling the 2D Slice:

// A loop iterates four times (for i := 0; i < 4; i++), creating a new slice arr with elements [1, 1, 1, 1] on each iteration.
// It appends this arr slice to a, adding one row at a time.
// Printing the 2D Slice:

// The second loop (for _, v := range a) iterates over each row in a, printing it out. This results in each row being displayed on a new line.

// [1 1 1 1]
// [1 1 1 1]
// [1 1 1 1]
// [1 1 1 1]
