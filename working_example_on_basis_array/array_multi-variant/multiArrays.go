package main

import (
	"fmt"
)

func main() {
	//declared as a 2D array with 3 rows and 4 columns. Each element is initialized with integer values, organized in nested braces for clarity.
	multiArray := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(multiArray)
	fmt.Println(multiArray[1][2]) // Accesses the element at row 1, column 2
	fmt.Println(multiArray[2][3]) // Accesses the element at row 2, column 3
}

// [[1 2 3 4] [5 6 7 8] [9 10 11 12]]
// 7
// 12
