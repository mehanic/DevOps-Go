package main

import "fmt"

// z is passed by value, so we are not passing the memory address, we're passing
// the value (value can't change!)
func zero(z int) {
	z = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}
