package main

import "fmt"

func main2() {
	fmt.Println("This is the implementation of closure stored in a variable")

	getMinMax := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	numbers := []int{5, 12, 42, 2, 12, 11}
	var min, max = getMinMax(numbers)
	fmt.Printf("Available numbers: %v\n Smallest number: %v\n Largest number: %v\n", numbers, min, max)
}
