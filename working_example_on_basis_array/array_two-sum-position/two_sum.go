package main

import (
	"fmt"
)

// twoSum function finds two indices of nums that add up to the target.
func twoSum(nums []int, target int) []int {
	dic := make(map[int]int) // map to store the needed difference for each number
	for index, num := range nums {
		if i, ok := dic[num]; ok { // check if num is in map
			return []int{i, index} // return indices if target pair is found
		}
		dic[target-num] = index // store the difference and current index
	}
	return []int{} // return empty if no solution found
}

func main() {
	// Example input
	nums := []int{2, 7, 11, 15}
	target := 9

	// Call twoSum and print result
	result := twoSum(nums, target)
	if len(result) > 0 {
		fmt.Printf("Indices: %d and %d\n", result[0], result[1])
	} else {
		fmt.Println("No two numbers add up to the target.")
	}
}

//Indices: 0 and 1
