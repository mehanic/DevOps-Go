package main

import (
	"fmt"
	"os"
)

func main() {
	var name string

	// assign a new value to the string variable below
	name = os.Args[1]
	fmt.Println("Hello great", name, "!")

	// changes the name, declares the age with 85
	name, age := "gandalf", 85

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}

// go run exec_input-args1/main.go arg0
// Hello great arg0 !
// My name is gandalf
// My age is 85
// BTW, you shall pass!
