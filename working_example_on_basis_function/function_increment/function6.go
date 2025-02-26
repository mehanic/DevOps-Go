package main

import "fmt"

func main() {
	var a = 8
	fmt.Println("a before: ", a)
	increment(a)
	fmt.Println("a after: ", a)
}
func increment(x int) {

	fmt.Println("x before: ", x)
	x = x + 20
	fmt.Println("x after: ", x)
}

// a before:  8
// x before:  8
// x after:  28
// a after:  8