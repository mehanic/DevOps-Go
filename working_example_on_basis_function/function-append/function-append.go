package main

import (
	"fmt"
)

func eggs(someParameter []interface{}) {
	someParameter = append(someParameter, "Hello")
	// This is a demonstration of appending within the function.
	// However, this won't modify the original slice since slices are reference types.
	// To modify the original slice, you need to use pointers or return the modified slice.
}

func main() {
	spam := []interface{}{1, 2, 3}
	eggs(spam)
	fmt.Println(spam)
}

