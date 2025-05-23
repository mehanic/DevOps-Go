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
	mini := arr[0][0]
	for _, v := range arr {
		fmt.Println(v)
		for _, k := range v {
			fmt.Println(k)
			if k < mini {
				mini = k
			}
		}
	}
	for i, v := range arr {
		for j, k := range v {
			if k == mini {
				fmt.Println(mini, i+1, j+1)
			}
		}
	}
}
