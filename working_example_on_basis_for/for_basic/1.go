package main

import "fmt"

func main() {
	n := 5
	numbers := [10]int{}
	fmt.Println(numbers)
	for i := 0; i < n; i++ {
		numbers[i] = 5
	}
	fmt.Println(numbers)
}
