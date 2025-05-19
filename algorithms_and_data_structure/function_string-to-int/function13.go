package main

import (
	"fmt"
)

func display() {
	fmt.Println("Welcome to functions concept")
	fmt.Println("Simple way to define your function")
}

func main() {
	// Calling the display function
	display()

	// Using the len equivalent in Go (for strings)
	str := "hi"
	fmt.Println(len(str))

	// Variable declaration and printing memory address using the & operator
	x := 40
	fmt.Printf("Address of x: %p\n", &x)

	// Working with tuples (Go uses slices or arrays instead)
	tuple := []int{5, 6}
	fmt.Println("Length of tuple:", len(tuple))
	fmt.Println("Minimum value:", min(tuple))
	fmt.Println("Maximum value:", max(tuple))

	// Converting string to integer
	numStr := "5"
	num, err := stringToInt(numStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println(num)
	}
}

// Helper function to get the minimum of an integer slice (Go has no built-in min/max for slices)
func min(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	minValue := slice[0]
	for _, val := range slice {
		if val < minValue {
			minValue = val
		}
	}
	return minValue
}

// Helper function to get the maximum of an integer slice
func max(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	maxValue := slice[0]
	for _, val := range slice {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

// Helper function to convert string to integer
func stringToInt(s string) (int, error) {
	var num int
	_, err := fmt.Sscanf(s, "%d", &num)
	return num, err
}

// Welcome to functions concept
// Simple way to define your function
// 2
// Address of x: 0xc00011a040
// Length of tuple: 2
// Minimum value: 5
// Maximum value: 6
// 5
