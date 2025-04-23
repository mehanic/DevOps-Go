package defers

import (
	"fmt"
)

/*

When weyou attempt to access an index beyond the length of a slice or the capacity of an array,
the Go runtime will generate a panic.

*/

func OutofBound() {
	capitals := []string{"Istanbul", "Berlin", "Vienna"}

	fmt.Println("Out of Bound Panic example")

	for capital := range capitals {
		fmt.Println(capital)
	}

	fmt.Println(capitals[3])
}
