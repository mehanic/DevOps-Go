package main

import (
	"fmt"
	"math/rand"
)

func CreateTwoDimensionalArray(rows, columns, randomRange int) [][]int {
	arr := [][]int{}
	//code
	for i := 0; i < rows; i++ {
		a := []int{}
		for j := 0; j < columns; j++ {
			k := rand.Intn(randomRange)
			a = append(a, k)
		}
		arr = append(arr, a)
	}
	return arr
}

func main() {
	arr := CreateTwoDimensionalArray(10, 20, 15)
	for _, v := range arr {
		fmt.Println(v)
	}
	for i := 0; i < len(arr); i++ {
		shetchik := 0
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 5 {
				shetchik += 1
			}
		}
		if shetchik >= 3 {
			fmt.Println("Повторение в массиве = ", shetchik, "", "Строка = ", i)
		}
	}
}

// [1 4 12 0 5 5 5 4 12 0 13 3 1 14 0 10 5 3 8 1]
// [6 8 4 5 6 1 14 4 10 5 12 1 5 9 13 8 5 13 6 1]
// [8 6 7 1 8 8 11 8 9 5 2 11 8 8 10 1 2 8 8 12]
// [13 6 10 5 4 3 9 13 3 3 13 12 5 12 5 5 13 9 13 12]
// [2 0 11 9 6 2 11 10 6 14 7 11 3 10 3 6 2 5 10 14]
// [8 2 12 3 7 8 9 1 14 6 13 13 6 3 9 13 13 1 13 1]
// [6 0 14 11 2 12 8 14 6 2 0 10 4 6 6 3 11 7 10 3]
// [0 13 9 3 9 13 7 4 0 5 5 4 6 6 5 0 8 12 14 13]
// [0 3 9 0 12 9 3 4 1 2 11 4 11 1 7 10 7 13 3 3]
// [6 5 14 7 2 5 6 9 9 12 14 12 11 13 3 1 9 1 2 0]
// Повторение в массиве =  4  Строка =  0
// Повторение в массиве =  4  Строка =  1
// Повторение в массиве =  4  Строка =  3
// Повторение в массиве =  3  Строка =  7
