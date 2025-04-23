package main

import "fmt"
//Brute Force
func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ { // Outer loop (i)
		firstNum := array[i]                  // First number in the pair
		for j := i + 1; j < len(array); j++ { // Inner loop (j) to compare with subsequent numbers
			secondNum := array[j]             // Second number in the pair
			if firstNum+secondNum == target { // Check if the sum equals the target
				return []int{firstNum, secondNum} // If yes, return the pair
			}
		}
	}
	return []int{} // If no pair is found, return an empty array
}

func main() {
	array := []int{2, 7, 11, 15}
	target := 9
	result := TwoNumberSum(array, target)
	fmt.Println(result, "this is result of array := []int{2, 7, 11, 15}") // Output: [2, 7] Explanation: 2 + 7 = 9
	main1()
	main2()
	main3()
	main4()
}

func main1() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	result := TwoNumberSum(array, target)
	fmt.Println(result, "this is result of array := []int{3, 5, -4, 8, 11, 1, -1, 6}") // Output: [-1, 11] Explanation: -1 + 11 = 10
}

func main2() {
	array := []int{1, 2, 3, 4, 5, 6}
	target := 8
	result := TwoNumberSum(array, target)
	fmt.Println(result, "this is result of array := []int{1, 2, 3, 4, 5, 6}") // Output: [2, 6] Explanation: 2 + 6 = 8
}

func main3() {
	array := []int{1, 3, 5, 7}
	target := 2
	result := TwoNumberSum(array, target)
	fmt.Println(result, "this is result of array := []int{1, 3, 5, 7}") // Output: [] Explanation: No two numbers sum to 2.
}

func main4() {
	array := []int{4, 4, 3, 7, 8, 9}
	target := 11
	result := TwoNumberSum(array, target)
	fmt.Println(result, "this is result of array := []int{4, 4, 3, 7, 8, 9}") // Output: [3, 8] (or another valid pair) Explanation: 3 + 8 = 11 (The function stops at the first valid pair).
}
