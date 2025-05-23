package main

import "fmt"

func main() {
	numbers := [...]int{10, 8, 1, 2, 3, 45, 12, 20}
	n := len(numbers)
	sumi := 0
	for i := 0; i < n; i++ {
		sumi = sumi + numbers[i]
		fmt.Println(sumi)
	}
	// n := len(numbers) — вычисление длины numbers
	fmt.Println(sumi / n)
}

// 12
