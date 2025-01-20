package main

import "fmt"

func main() {
	numbers := []int{}
	n := 100
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			numbers = append(numbers, i)
		}
	}
	indexes := []int{2, 6, 3}
	for _, v := range indexes {
		for j, value := range numbers {
			if v-1 == j {
				fmt.Print(value, " ")
			}
		}
	}
	//0 2 4 6 8 10 12 ...
	//2 6 3
	//2 10 4
}
