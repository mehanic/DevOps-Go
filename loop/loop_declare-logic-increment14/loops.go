package main

import "fmt"

func SumLoop() int {
	result := 0
	for i := 0; i < 10; i++ {
		result += i
	}
	return result

}

func SumArray(arr []int) int {
	var result int
	for _, num := range arr {
		result += num
	}
	return result
}

func ListElements(elements []string) {
	for index, value := range elements {
		fmt.Println(index, value)
	}

}


func main() {
	// Example 1: Use SumLoop to calculate the sum of numbers from 0 to 9
	sumResult := SumLoop()
	fmt.Println("The sum of numbers from 0 to 9 is:", sumResult)

	// Example 2: Use SumArray to calculate the sum of elements in an array
	arr := []int{1, 2, 3, 4, 5}
	arraySumResult := SumArray(arr)
	fmt.Println("The sum of the array elements is:", arraySumResult)

	// Example 3: Use ListElements to list all elements with their indices
	elements := []string{"Go", "Python", "JavaScript", "Java", "C++"}
	fmt.Println("Listing elements with indices:")
	ListElements(elements)
}

// The sum of numbers from 0 to 9 is: 45
// The sum of the array elements is: 15
// Listing elements with indices:
// 0 Go
// 1 Python
// 2 JavaScript
// 3 Java
// 4 C++
