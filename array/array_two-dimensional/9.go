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
