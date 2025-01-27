package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	different types of variables
	var myName string = "Timothy Unkert"
	var age int = 46
	var isTeacher bool = true
	var pi float64 = 3.14
	fmt.Println("Hi, my name is " + myName + " and I am " + strconv.Itoa(age) + "years old.")
	// I use the strconv package to convert an integer to a string here
	fmt.Println(isTeacher)
	fmt.Println(pi)
}

// Hi, my name is Timothy Unkert and I am 46years old.
// true
// 3.14
