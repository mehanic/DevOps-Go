package main

import "fmt"

// Function to demonstrate finding the largest number in a 1D array
func Array1() {
	numbers := [5]int{233, 24, 56, 87, 23}
	fmt.Println(numbers)
	fmt.Println("-------------------")

	// Initialize maxValue with the first element of the array
	maxValue := numbers[0]

	// Iterate through the array to find the maximum value
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		if numbers[i] > maxValue {
			maxValue = numbers[i]
		}
	}
	fmt.Println(maxValue)
	fmt.Println("-------------------")
}

func Array2() {
	var numbers [2][3]int
	numbers[0][0] = 12
	numbers[0][1] = 31
	numbers[0][2] = 98
	numbers[1][0] = 23
	numbers[1][1] = 54
	numbers[1][2] = 13
	fmt.Println(numbers)
}

func main() {
	//	Array1()
	Array2()
}

// [233 24 56 87 23]
// -------------------
// 233
// 24
// 56
// 87
// 23
// 233
// -------------------
// [[12 31 98] [23 54 13]]
