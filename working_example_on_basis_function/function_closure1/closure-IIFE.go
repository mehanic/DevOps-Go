package main

import "fmt"

func main() {
	fmt.Println("This is the implementation of the closure Immediately Invoked Function Expression (IIFE)")
	// Example of applying IIFE for filtering data in an array
	numbers := []int{5, 1, 35, 6, 3, 3, 1, 312}
	newNumbers := func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)
	fmt.Println("Original Numbers =", numbers)
	fmt.Println("Filtered Numbers =", newNumbers)
}
