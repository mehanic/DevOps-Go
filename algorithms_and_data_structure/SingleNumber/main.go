package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
    unique := 0
    for _, num := range nums {
        unique ^= num  // XOR all numbers
    }
    return unique
}

func main() {
    fmt.Println(singleNumber([]int{3, 2, 3}))       // Output: 2
    fmt.Println(singleNumber([]int{7, 6, 6, 7, 8})) // Output: 8
    fmt.Println(singleNumber([]int{1, 1, 2, 2, 5})) // Output: 5
    fmt.Println(singleNumber([]int{9}))            // Output: 9 (single element case)
}
