package main

import "fmt"

// Function spam prints the global eggs and then declares a local variable
func spam() {
	fmt.Println(eggs) // This will print the global variable
	eggs := "spam local"
	fmt.Println(eggs) // This will print the local variable
}

var eggs = "global"

func main() {
	// Call spam function
	spam()
}

