package main

import (
	"fmt"
)

func jumpToTheEnd(nums []int) bool {
	destination := len(nums) - 1
	// Проходим массив в обратном порядке
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= destination {
			destination = i
		}
	}
	// Если дошли до начала, значит можно прыгнуть до конца
	return destination == 0
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jumpToTheEnd(nums)) // Output: true

	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(jumpToTheEnd(nums)) // Output: false
}
