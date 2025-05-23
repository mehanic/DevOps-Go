package main

import (
	"fmt"
)

func find4(num int, nums ...int) {
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
	nums := []int{89, 90, 95}
	find4(89, nums...)
}

// type of nums is []int
// 89 found at index 0 in [89 90 95]

