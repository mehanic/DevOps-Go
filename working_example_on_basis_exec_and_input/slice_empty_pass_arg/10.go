package main

import "fmt"

func main() {
	//arr := [][]int{
	//	{25, 1},
	//	{70, 1},
	//	{100, 0},
	//	{3, 1},
	//}
	arr := [][]int{}
	var n int
	fmt.Scanf("%d \n", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d \n", &a, &b)
		nestedArr := []int{a, b}
		arr = append(arr, nestedArr)
	}
	fmt.Println(arr)
}

// Dynamic Array Creation Based on User Input:

// arr := [][]int{} initializes an empty 2D slice, preparing it to hold integer sub-arrays.
// var n int defines an integer variable n to represent the number of sub-arrays the user wishes to add.
// fmt.Scanf("%d \n", &n) prompts the user to input a single integer for n, specifying how many rows (sub-arrays) will be entered.
// Loop to Populate Array:

// The for loop iterates n times. In each iteration:
// var a, b int declares two integer variables, a and b, to store values for each sub-array.
// fmt.Scanf("%d %d \n", &a, &b) reads two integers from the user.
// nestedArr := []int{a, b} creates a sub-array with the two integers.
// arr = append(arr, nestedArr) appends the sub-array to the main 2D array arr
