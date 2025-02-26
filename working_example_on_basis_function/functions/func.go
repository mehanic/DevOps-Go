package functions

import "fmt"

func CreateFunction() {

	/*
		A function is a block of statements that can be used repeatedly in a program.
		A function will not execute automatically when a page loads.
		A function will be executed by a call to the function.

	*/

	fmt.Println("This is a main function")

	fmt.Println("What's your name?")

	var name string

	fmt.Scanln(&name)

	user_name := Greetings(name)

	fmt.Println("\nUser Name", user_name)

}

func Greetings(name string) string { // We give arguments with this way. We also return a string.

	fmt.Printf("Welcome %v , We learn golang here!!!", name)

	return name // We return the name for future usage

}
