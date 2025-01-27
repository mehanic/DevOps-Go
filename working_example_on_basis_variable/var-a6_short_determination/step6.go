package main

import "fmt"

func main() {

	box := "Hello"
	fmt.Println(box) //bunda faqat variableni o'zi bo'ladi !

	// fmt.Printf("", box ) --> bunda 2 ta pqrametr oladi "" va box (variable)
	fmt.Printf("Qalle: %s \n", box)
	fmt.Printf("Salom: %v \n", box)

	// Printing multiple values with %v
	fmt.Printf("%v %v %v %v \n", 11, 22.5, "Alxan", true) // VALUES
	// Using %T to print the type of each variable
	x, y, z, b := 1, 77.6, "World", false
	fmt.Printf("1: %T\n2: %T\n3: %T\n4: %T\n", x, y, z, b)
}

// Hello
// Qalle: Hello
// Salom: Hello
// 11 22.5 Alxan true
// 1: int
// 2: float64
// 3: string
// 4: bool
