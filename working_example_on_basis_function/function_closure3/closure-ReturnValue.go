package main

import "fmt"

func findMax(number []int, max int) (int, func() []int) {
	var res []int
	for _, e := range number {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

func main() {
	fmt.Println("Applying Closure with return value")

	maxNumber := 5
	dataNumbers := []int{2, 6, 3, 3, 36, 2, 7, 3, 4, 2}
	howMany, getNumber := findMax(dataNumbers, maxNumber)
	resultNumbers := getNumber()

	fmt.Println("The numbers are: ", dataNumbers)
	fmt.Println("Found: ", howMany, "numbers below the maximum number")
	fmt.Println("Numbers below the maximum are: ", resultNumbers)
}
