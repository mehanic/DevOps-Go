package main

import "fmt"

func main() {
	add3(1, 2, 3)       // sum = 6
	add3(1, 2, 3, 4)    // sum = 10
	add3(5, 6, 7, 2, 3) // sum = 23
}

func add3(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

// sum =  6
// sum =  10
// sum =  23

