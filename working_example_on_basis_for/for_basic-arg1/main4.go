package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	numbers := [][]int{}
	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < m; j++ {
			var d int
			fmt.Scan(&d)
			temp = append(temp, d)
		}
		numbers = append(numbers, temp)
	}
	fmt.Println(numbers)
}
