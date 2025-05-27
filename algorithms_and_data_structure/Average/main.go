package main

import (
	"fmt"
)

func main() {
	av := avg(12, 23, 45, 45)
	fmt.Println(av)
}

func avg(num ...float64) float64 {
	fmt.Println(num)
	fmt.Printf("%T \n", num)
	total := 0.0
	for _, value := range num {
		total += value
	}

	return total / float64(len(num))
}

// avg function is called a variadic function since it take a slice of float64 values

// [12 23 45 45]
// []float64
// 31.25
