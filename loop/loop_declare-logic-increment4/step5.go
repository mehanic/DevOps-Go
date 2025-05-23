package main

import "fmt"

func main() {
	// Iterate from 1 to 100
	for n := 1; n <= 100; n++ {
		// Skip even numbers
		if n%2 == 0 {
			continue
		}
		// Print odd numbers
		fmt.Println(n)
	}
}


// 89
// 91
// 93
// 95
// 97
// 99
