package main

import "fmt"

func main() {
	arr := [][]int{
		{5, 5, 2, 2},
		{4, 5, 5, 5},
		{5, 1, 23, 4},
		{5, 5, 5, 5},
	}
	for i, v := range arr {
		count := 0
		for _, k := range v {
			if k == 5 {
				count += 1
			}
		}
		if count >= 3 {
			fmt.Println(i)
		}
	}
	main1()
}

func findRowsWithValue(arr [][]int, val int, countThreshold int) []int {
	var result []int

	for i, row := range arr {
		count := 0
		for _, num := range row {
			if num == val {
				count++
			}
		}
		if count >= countThreshold {
			result = append(result, i)
		}
	}

	return result
}

func main1() {
	arr := [][]int{
		{5, 5, 2, 2},
		{4, 5, 5, 5},
		{5, 1, 23, 4},
		{5, 5, 5, 5},
	}

	indices := findRowsWithValue(arr, 5, 3)
	fmt.Println(indices) // [1 3]
}

