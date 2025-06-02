package main

import "fmt"

func Map(input []int, f func(int) int) []int {
	result := make([]int, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func main() {
	square := func(n int) int { return n * n }
	squared := Map([]int{1, 2, 3}, square)
	fmt.Println(squared)
	// => [1 4 9]

}
