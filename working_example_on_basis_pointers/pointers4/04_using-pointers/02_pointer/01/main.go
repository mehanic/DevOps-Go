package main

import "fmt"

// zero now takes in a pointer
func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}
