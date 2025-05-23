package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func CreateTwoDimensionalArray(rows, columns, randomRange int) [][]int {
	arr := [][]int{}
	// code
	for i := 0; i < rows; i++ {
		a := []int{}
		for j := 0; j < columns; j++ {
			// create random number
			k := rand.Intn(randomRange)
			// add elements to the nested array
			a = append(a, k)
		}
		// add array into arr array
		arr = append(arr, a)
	}
	return arr
}

func CreateArray(n, randomRange int) []int {
	//	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < n; i++ {
		k := rand.Intn(randomRange)
		arr = append(arr, k)
	}
	return arr
}

func main() {
	arr := CreateTwoDimensionalArray(3, 4, 20)
	for _, v := range arr {
		fmt.Println(v)
	}
}

// CreateTwoDimensionalArray Function:

// Parameters: rows (number of rows), columns (number of columns), and randomRange (upper limit for random numbers).
// The function initializes an empty 2D slice arr, iterating rows times to create each row.
// For each row, it generates a nested slice a that fills up with columns random integers between 0 and randomRange - 1.
// After filling the row, it appends it to the main 2D array arr.
// Finally, it returns arr, which is the complete 2D array filled with random numbers.
// CreateArray Function:

// Parameters: n (number of elements) and randomRange (upper limit for random values).
// rand.Seed(time.Now().UnixNano()) is used to ensure randomness in each program run by seeding with the current time.
// It creates an empty slice arr and fills it with n random integers within the specified range.
// Returns the 1D array arr.
// main Function:

// Calls CreateTwoDimensionalArray with 3 rows, 4 columns, and 20 as the random range.
// Loops through each row in the returned 2D array and prints it, displaying the randomly generated 2D array.

// [8 2 16 2]
// [17 4 8 4]
// [12 19 2 16]
