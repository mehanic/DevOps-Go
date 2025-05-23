package main

import "fmt"

// Function spam prints a local variable
func spam() {
	eggs := "spam local"
	fmt.Println(eggs) // prints 'spam local'
}

// Function bacon prints a local variable, calls spam, and prints the local variable again
func bacon() {
	eggs := "bacon local"
	fmt.Println(eggs) // prints 'bacon local'
	spam()
	fmt.Println(eggs) // prints 'bacon local'
}

func main() {
	eggs := "global"
	bacon()
	fmt.Println(eggs) // prints 'global'
}

