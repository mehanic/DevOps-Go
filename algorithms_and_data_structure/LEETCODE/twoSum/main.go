package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) // map[num] = index

	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1, 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0, 1]
	nums := []int{1, 2, 3}
	target := 10
	fmt.Println(twoSum(nums, target)) // => nil
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum3(nums []int, target int) []int {
    	for i := 0; i < len(nums)-1; i++ { // Outer loop (i)
		firstNum := nums[i]                  // First number in the pair
		for j := i + 1; j < len(nums); j++ { // Inner loop (j) to compare with subsequent numbers
			secondNum := nums[j]             // Second number in the pair
			if firstNum+secondNum == target { // Check if the sum equals the target
				return []int{i, j} // If yes, return the pair
			}
		}
	}
	return []int{} // If no pair is found, return an empty array

}