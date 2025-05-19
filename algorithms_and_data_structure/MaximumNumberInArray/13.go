package main

import "fmt"

func main() {
	numbers := [...]int{10, 8, 1, 2, 3, 45, 12, 20}
	n := len(numbers)
	maxi := numbers[0]
	for i := 0; i < n; i++ {
		if numbers[i] > maxi {
			maxi = numbers[i]
		}
	}
	fmt.Println(maxi)
}

// 45
