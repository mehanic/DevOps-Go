package main

import "fmt"

func main() {
	var mySlice []int // nil slice

	// Append will
	mySlice = append(mySlice, 1, 2, 3, 4, 5)

	firstElement := mySlice[0]
	fmt.Println("First element:", firstElement)

	subset := mySlice[1:4]
	fmt.Println(subset)

	subset = mySlice[1:]
	fmt.Println(subset)

	subset = mySlice[0 : len(mySlice)-1]
	fmt.Println(subset)

	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 4)

	// Create a unique copy in memory
	copy(slice2, slice1)

	// Changing one should not affect the other
	slice2[3] = 99
	fmt.Println(slice1)
	fmt.Println(slice2)
}
