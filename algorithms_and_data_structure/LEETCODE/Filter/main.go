package main

import "fmt"

func Filter(input []int, predicate func(int) bool) []int {
	result := make([]int, 0)
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {

	isEven := func(n int) bool { return n%2 == 0 }
	filtered := Filter([]int{1, 2, 3, 4}, isEven)
	fmt.Println(filtered)
	// => [2 4]

}
