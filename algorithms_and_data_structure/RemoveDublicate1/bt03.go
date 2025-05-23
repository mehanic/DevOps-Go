package main

import (
	"fmt"
	"os"
)

func main() {
	inputarray(7)
}

func inputarray(n int) {
	fmt.Println("Exercise 3: Write a function to remove duplicate elements from an array")
	fmt.Println("Example: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]")
	fmt.Println("---------------------------------------------------------------------")
	input := []int{}
	for i := 0; i < n; i++ {
		fmt.Printf("Enter array[%d] = ", i)
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Println("Invalid input!")
			return
		}
		input = append(input, number)
	}
	fmt.Println("The array you entered contains the following elements: ", input)
	removeDuplicate(input)
}

func removeDuplicate(a []int) (result []int) {
	keys := map[int]bool{} // You could also write: keys := make(map[int]bool) but it's more verbose
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	fmt.Printf("Array after removing duplicate elements: %d", result)
	fmt.Println()
	return
}
