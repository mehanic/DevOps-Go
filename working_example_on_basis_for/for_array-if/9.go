package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 7, 8, 12323}
	n := len(numbers)
	fmt.Println(numbers)
	for i := 0; i < n; i++ {
		if i != 2 {
			numbers[i] = numbers[i] * 2
		}

	}
	fmt.Println(numbers)
}

// [1 2 3 4 5 7 8 12323]
// [2 4 3 8 10 14 16 24646]