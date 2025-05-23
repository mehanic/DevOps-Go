package main

import "fmt"

func main() {
	arr := [][]int{
		{-19, 6, 3, 18, -12, -3, 24},
		{-4, 15, -6, 19, -15, -1, 4},
		{6, -9, -12, 23, -3, 3, -11},
		{5, 0, -11, -4, -19, -6, 1},
		{17, 20, -1, 6, 17, -1, 15},
	}
	rows := len(arr)
	col := len(arr[0])
	for i := 0; i < col; i++ {
		maxi := arr[0][i]
		for j := 0; j < rows; j++ {
			if arr[j][i] > maxi {
				maxi = arr[j][i]
			}
		}
		fmt.Println(maxi)
	}
}
