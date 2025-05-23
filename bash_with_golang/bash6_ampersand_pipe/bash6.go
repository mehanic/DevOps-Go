package main

import (
	"fmt"
)

func main() {
	myName := "John"
	name1 := "Bob"
	name2 := "Jane"
	name3 := "Sue"
	name4 := "Kate"

	if myName == "Ron" {
		fmt.Println("Ron is home from vacation")
	} else if myName != name1 && myName != name2 && myName == "John" {
		fmt.Println("John is home after some unnecessary AND logic")
	} else if myName == name3 || myName == name4 {
		fmt.Println("Looks like one of the ladies are home")
	} else {
		fmt.Println("Who is this stranger?")
	}
}

