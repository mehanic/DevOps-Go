package main

import (
	"fmt"
	"strconv"
)

func reverseDigits(n int) []int {
	str := strconv.Itoa(n)
	var result []int
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, int(str[i]-'0')) // Append digits in reverse order
	}
	return result
}

func main() {
	fmt.Println(reverseDigits(35231)) // Output: [1 3 2 5 3]
}

