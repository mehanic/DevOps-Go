package main

import (
	"fmt"
)

func main() {
	// Define constants and variables
	PI := 3.14
	VAR_A := 10
	VAR_B := VAR_A
	VAR_C := VAR_B

	fmt.Println("Lets print 3 variables:")
	fmt.Println(VAR_A)
	fmt.Println(VAR_B)
	fmt.Println(VAR_C)

	fmt.Println("We know this will break:")
	// In Go, accessing an undefined variable will cause a compile-time error
	// So we'll simulate this by printing a message instead
	fmt.Println("0. The value of PI is [undefined]") // Go doesn't allow PIabc to be used.

	fmt.Println("And these will work:")
	fmt.Printf("1. The value of PI is %f\n", PI)
	fmt.Printf("2. The value of PI is %.2f\n", PI)
	fmt.Printf("3. The value of PI is %f\n", PI)

	fmt.Println("And we can make a new string")
	STR_A := "Bob"
	STR_B := "Jane"
	fmt.Printf("%s + %s equals Bob + Jane\n", STR_A, STR_B)

	STR_C := STR_A + " + " + STR_B
	fmt.Printf("%s is the same as Bob + Jane too!\n", STR_C)
	fmt.Printf("%s + %f\n", STR_C, PI)

	// Exit with a success status (not required in Go, as the program exits automatically)
}

