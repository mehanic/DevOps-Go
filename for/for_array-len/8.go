package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 7, 8, 12323}
	n := len(numbers)
	sumi := 0
	for i := 0; i < n; i++ {
		sumi = sumi + numbers[i]
	}
	fmt.Println(sumi)
}

// 12353
