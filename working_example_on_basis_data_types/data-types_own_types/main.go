package main

import (
	"fmt"
	"strings"
)

type mile float64
type kilometer float64

func main() {
	// Initial conversions and demonstrations
	m1 := mile(10)
	k1 := toKilometer(m1)
	fmt.Println("10 miles in kilometers:", k1)

	k2 := kilometer(10)
	m2 := toMile(k2)
	fmt.Println("10 kilometers in miles:", m2)

	// Demonstrate compatible type operations
	f1 := float64(4.4)
	fmt.Printf("%T, %v\n", m1+mile(f1), m1+mile(f1))
	fmt.Printf("%T, %v\n", float64(m1)+f1, float64(m1)+f1)

	// Demonstrate adding two miles
	var m3 mile = 3.2
	fmt.Printf("m3: %T, %v\n", m3, m3)
	m4 := mile(4.6)
	fmt.Println("Sum of two miles (m3 + m4):", m3+m4)
	fmt.Printf("Product of two miles (m3 * m4): %.2f\n", m3*m4)

	// Print string in uppercase
	s1 := "arin"
	fmt.Println("Uppercase of 'arin':", strings.ToUpper(s1))
}

// Function to convert miles to kilometers
func toKilometer(m mile) kilometer {
	return kilometer(m * 1.6)
}

// Function to convert kilometers to miles
func toMile(k kilometer) mile {
	return mile(k * 0.62)
}

// 10 miles in kilometers: 16
// 10 kilometers in miles: 6.2
// main.mile, 14.4
// float64, 14.4
// m3: main.mile, 3.2
// Sum of two miles (m3 + m4): 7.8
// Product of two miles (m3 * m4): 14.72
// Uppercase of 'arin': ARIN
