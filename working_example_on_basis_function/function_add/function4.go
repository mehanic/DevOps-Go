package main

import "fmt"

func main() {
	add(4, 5)  // x + y = 9
	add(20, 6) // x + y = 26
}

func add(x int, y int) {
	var z = x + y
	fmt.Println("x + y = ", z)
}

// x + y =  9
// x + y =  26
