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
}
