package main

import "fmt"

func main() {
	result := sum(20, 10, 50)
	fmt.Println("result is = ", result)
}

func sum(nums ...int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

// result is =  80
