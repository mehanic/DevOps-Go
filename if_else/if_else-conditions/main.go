package main

import "fmt"

func main() {
	// Basic condition
	if true {
		fmt.Println("this should be printed")
	}

	if false {
		fmt.Println("this will never be printed")
	}

	// Multiple conditions
	condition01 := 1 == 1
	condition02 := 1 == 2

	if condition01 {
		fmt.Println("condition 1 is met")
	} else if condition02 {
		fmt.Println("condition 2 is met")
	} else {
		fmt.Println("none of the conditions are met")
	}

	// Ternary-like conditional
	defaultValue := 2
	if condition01 {
		defaultValue = 1
	}

	fmt.Println(defaultValue)
}


// this should be printed
// condition 1 is met
// 1
