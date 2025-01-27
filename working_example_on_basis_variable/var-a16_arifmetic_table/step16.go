// package main

// import "fmt"

// func main() {

// 	var one int = 1
// 	var two int = 2
// 	var three int = 3
// 	var four int = 4
// 	var five int = 5
// 	var six int = 6
// 	var seven int = 7
// 	var eight int = 8
// 	var nine int = 9
// 	var ten int = 10
// 	var multiply string = "*"
// 	var equals string = "="

// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, one, equals, one*one)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, two, equals, one*two)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, three, equals, one*three)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, four, equals, one*four)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, five, equals, one*five)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, six, equals, one*six)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, seven, equals, one*seven)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, eight, equals, one*eight)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, nine, equals, one*nine)
// 	fmt.Printf("%d %s %d %s %d  \t", one, multiply, ten, equals, one*ten)

// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, one, equals, two*one)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, two, equals, two*two)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, three, equals, two*three)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, four, equals, two*four)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, five, equals, two*five)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, six, equals, two*six)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, seven, equals, two*seven)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, eight, equals, two*eight)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, nine, equals, two*nine)
// 	fmt.Printf("%d %s %d %s %d  \t", two, multiply, ten, equals, two*ten)

// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, one, equals, three*one)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, two, equals, three*two)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, three, equals, three*three)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, four, equals, three*four)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, five, equals, three*five)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, six, equals, three*six)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, seven, equals, three*seven)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, eight, equals, three*eight)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, nine, equals, three*nine)
// 	fmt.Printf("%d %s %d %s %d   \t", three, multiply, ten, equals, three*ten)

// 	// This code continues similarly for four, five, six, seven, eight, nine, and ten.
// 	// Essentially, it's a program to generate the multiplication table from 1 to 10.

// }

package main

import (
	"fmt"
	//"strconv"
)

func main() {
	kpy := "*"
	shox := "="

	// Outer loop for each number from 1 to 10
	for i := 1; i <= 10; i++ {
		// Inner loop for multiplication with each number from 1 to 10
		for j := 1; j <= 10; j++ {
			// Printing each multiplication result with the format
			fmt.Printf("%d %s %d %s %d\t", i, kpy, j, shox, i*j)
		}
		fmt.Println() // New line after each row
	}
}
