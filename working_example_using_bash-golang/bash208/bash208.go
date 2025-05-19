package main

import "fmt"

// Declare a global variable
var eggs = "global"

// Function to modify the global variable
func spam() {
	eggs = "spam" // This modifies the global variable
}

func main() {
	// Call the spam function
	spam()
	// Print the global variable, which is now modified
	fmt.Println(eggs) // Output: spam
}

