package main

import "fmt"

func TwoSum(array []int, target int) []int {
	for i := 0; i < len(array) - 1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++{
		secondNum := array[j]
		if firstNum+secondNum == target {
			return []int{firstNum, secondNum}
		}
	}
}
  return []int{}

}

func main() {
	array := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(array, target)
	fmt.Println(result, "this is result of array := []int{2, 7, 11, 15}") // Output: [2, 7] Explanation: 2 + 7 = 9
}



