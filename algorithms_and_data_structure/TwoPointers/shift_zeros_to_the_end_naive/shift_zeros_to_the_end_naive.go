package main

import (
	"fmt"
)

func shiftZerosToTheEndNaive(nums []int) {
	n := len(nums)
	temp := make([]int, n) // Инициализируем массив нулями (по умолчанию)

	i := 0
	for _, num := range nums {
		if num != 0 {
			temp[i] = num
			i++
		}
	}

	for j := 0; j < n; j++ {
		nums[j] = temp[j]
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	shiftZerosToTheEndNaive(nums)
	fmt.Println("After shifting zeros:", nums) // Output: [1 3 12 0 0]
}
