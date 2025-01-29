package main

import "fmt"

func main() {
	twoArr := [][]int{
		{1, 2, 3, 5},
		{4, 5, 6, 0},
		{7, 8, 9, 1},
		{1, 2, 3, 5},
	}
	//for i := 0; i < len(twoArr); i++ {
	//	fmt.Println("1")
	//	for j := 0; j < len(twoArr[i]); j++ {
	//		fmt.Println("2")
	//	}
	//}

	// Loop with range:

	// The outer loop (for i, v := range twoArr) iterates over each row of twoArr.
	// The inner loop (for l, k := range v) iterates over each element in the current row.
	// i and l are indices, while v and k are values.
	// Sum Calculation:

	// Each element (k) in twoArr is added to the variable sumi, which accumulates the total sum of all elements.
	// Element Replacement:

	// After adding each element’s value to sumi, the code sets that element to 0 by updating twoArr[i][l].
	// Output:

	//fmt.Println(sumi) displays the sum of all original values in twoArr.
	//fmt.Println(twoArr) shows twoArr after all elements have been replaced with 0.
	sumi := 0
	for i, v := range twoArr {
		for l, k := range v {
			sumi += k
			twoArr[i][l] = 0
		}
	}
	fmt.Println(sumi)
	fmt.Println(twoArr)
}

// 62
// [[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
