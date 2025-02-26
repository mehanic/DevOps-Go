package main

import (
	"fmt"
	"reflect"
)

func display(args ...interface{}) {
	for _, each := range args {
		fmt.Println(reflect.TypeOf(each))
	}
}

func main() {
	// display() // Uncomment to test with no arguments
	// display(4) // Uncomment to test with a single argument
	display(4, 5, 67)
	fmt.Println("-------------")
	display("hi", 4.65)
}

// int
// int
// int
// -------------
// string
// float64
