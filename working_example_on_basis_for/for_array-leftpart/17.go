package main

import "fmt"

func main() {
	numbers := [...]int{10, 8, 1, 2, 3, 45, 12, 20}
	var leftPart [10]int
	n := len(numbers)
	j := 0
	for i := 0; i < n/2; i++ {
		// leftPart[i] = numbers[i]
		leftPart[j] = numbers[i]
		j++
	}
	fmt.Print(leftPart)
	fmt.Println("difference between slice and array")
	main2()
}

// [10 8 1 2 0 0 0 0 0 0]

func main2() {
	numbers := [...]int{10, 8, 1, 2, 3, 45, 12, 20}
	var leftPart []int
	n := len(numbers)
	for i := 0; i < n/2; i++ {
		leftPart = append(leftPart, numbers[i])
	}
	fmt.Print(leftPart)
}
