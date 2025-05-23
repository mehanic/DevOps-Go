package main

import (
	"fmt"
)

func main() {
	// Define the array
	array := []string{"Debian Linux", "Redhat Linux", "Ubuntu Linux"}

	// Get the number of elements in the array
	elements := len(array)

	// Loop through each element in the array
	for i := 0; i < elements; i++ {
		fmt.Println(array[i])
	}
}

