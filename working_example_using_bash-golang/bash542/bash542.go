package main

import (
	"fmt"
)

func main() {
	// Declare array with 3 elements
	array := [3]string{"Debian Linux", "Redhat Linux", "Ubuntu Linux"}

	// Get number of elements in the array
	elements := len(array)

	// Echo each element in the array
	for i := 0; i < elements; i++ {
		fmt.Println(array[i])
	}
}

