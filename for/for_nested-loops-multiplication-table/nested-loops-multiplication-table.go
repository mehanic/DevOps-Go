package main

import (
	"fmt"
)

// EXERCISE: Get the `max` from the user
//           And create the table according to that.

const max = 5

func main() {
	// print the header
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		// print the vertical header
		fmt.Printf("%5d", i)

		// print the cells
		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

// X    0    1    2    3    4    5
// 0    0    0    0    0    0    0
// 1    0    1    2    3    4    5
// 2    0    2    4    6    8   10
// 3    0    3    6    9   12   15
// 4    0    4    8   12   16   20
// 5    0    5   10   15   20   25
