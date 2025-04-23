package main

import "fmt"

func main() {
	twoArr := [][]int{
		{1, 2, 3, 5},
		{4, 5, 6, 0},
		{7, 8, 9, 1},
		{1, 2, 3, 5},
	}

	//For each row in twoArr, sumi is initialized to zero to store
	// the sum of elements in that row.
	//The inner loop iterates over each element in the row
	//to calculate the row’s total sum (sumi).
	//After the inner loop completes, sumi / len(twoArr[i]) is
	//printed as the average of the row’s elements.
	for i := 0; i < len(twoArr); i++ {
		sumi := 0
		fmt.Println(twoArr[i])
		for j := 0; j < len(twoArr[i]); j++ {
			sumi += twoArr[i][j] //// Accumulate the sum of elements in the row
		}
		fmt.Println(sumi / len(twoArr[i]))
	}
}

// [1 2 3 5]
// 2
// [4 5 6 0]
// 3
// [7 8 9 1]
// 6
// [1 2 3 5]
// 2
