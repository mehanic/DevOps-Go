package main

import "C"

import (
	"fmt"
	"math"
)

func main() {
	a := float64(math.Pi)

	fmt.Println(a)
	fmt.Println(C.float(a))
	fmt.Println(C.double(a))
	fmt.Println(C.double(C.float(a)) - C.double(a))
}

// 3.141592653589793
// 3.1415927
// 3.141592653589793
// 8.742278012618954e-08
