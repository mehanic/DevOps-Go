package main

import (
	"fmt"
)

func find1(num int, nums []int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func main() {
	find1(89, []int{89, 90, 95})
	find1(45, []int{56, 67, 45, 90, 109})
	find1(78, []int{38, 56, 98})
	find1(87, []int{})
}

// type of nums is []int
// 89 found at index 0 in [89 90 95]

// type of nums is []int
// 45 found at index 2 in [56 67 45 90 109]

// type of nums is []int
// 78 not found in  [38 56 98]

// type of nums is []int
// 87 not found in  []

