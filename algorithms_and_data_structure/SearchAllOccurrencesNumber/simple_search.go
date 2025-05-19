package main

import (
	"fmt"
)

// Function to search for a number in the list
func ListSearch(list []int, number int) (resultSearch []int) {
	var result []int
	for _, list_number := range list {
		if list_number == number {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	numbers := []int{3, 5, 7, 5, 9, 5, 2}
	target := 5

	result := ListSearch(numbers, target)
	fmt.Println("Occurrences of", target, ":", result) // Output: Occurrences of 5 : [5 5 5]
}
