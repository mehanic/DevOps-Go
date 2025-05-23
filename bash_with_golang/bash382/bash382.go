package main

import (
	"fmt"
)

func main() {
	// Variable declarations
	MY_NAME := "John"
	NAME_1 := "Bob"
	NAME_2 := "Jane"
	NAME_3 := "Sue"
	NAME_4 := "Kate"

	// Conditional logic
	if MY_NAME == "Ron" {
		fmt.Println("Ron is home from vacation")
	} else if MY_NAME != NAME_1 && MY_NAME != NAME_2 && MY_NAME == "John" {
		fmt.Println("John is home after some unnecessary AND logic")
	} else if MY_NAME == NAME_3 || MY_NAME == NAME_4 {
		fmt.Println("Looks like one of the ladies are home")
	} else {
		fmt.Println("Who is this stranger?")
	}
}

