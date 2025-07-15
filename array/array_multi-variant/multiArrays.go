package main

import (
	"fmt"
)

func main() {
	multiArray := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(multiArray)
	fmt.Println(multiArray[1][2]) // Accesses the element at row 1, column 2
	fmt.Println(multiArray[2][3]) // Accesses the element at row 2, column 3
}

// [[1 2 3 4] [5 6 7 8] [9 10 11 12]]
// 7
// 12
