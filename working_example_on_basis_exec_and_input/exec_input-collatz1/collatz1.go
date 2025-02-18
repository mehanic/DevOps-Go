package main

import (
	"fmt"
)

var result int

func collatz(number int) int {
	if number%2 == 0 {
		fmt.Println(number / 2)
		result = number / 2
		return result
	} else if number%2 == 1 {
		fmt.Println(3*number + 1)
		result = 3*number + 1
		return result
	}
	return result
}

func main() {
	fmt.Println("Enter number:")
	for {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("The value must be an integer!")
			continue
		}
		collatz(number)
		if result == 1 {
			break
		}
	}
}

// Enter number:
// 12
// 6
// 6
// 3
