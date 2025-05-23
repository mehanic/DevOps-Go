package main

import (
	"fmt"
	"math"
)

//Enter the three sides of a triangle from the keyboard.
//Calculate the area and perimeter
//of the triangle based on the given three sides.
//The area and perimeter of the triangle can be calculated using the following formulas.
//
//Perimeter of the triangle: P = a + b + c
//Area of the triangle: S = √(p(p-a)(p-b)(p-c)), where p = P/2

func main() {
	var a, b, c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	fmt.Println(s)
}

// 5
// 7
// 4
// 9.797958971132712
