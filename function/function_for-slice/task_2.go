package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var summ = 0
	for _, v := range slice {
		summ += v
	}

	fmt.Println(summ)
}

// 55
