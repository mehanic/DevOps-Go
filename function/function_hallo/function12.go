package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hallo Herr, %s\n", name)
}
func main() {
	name := "Peter"
	greet(name)
}

// Hallo Herr, Peter
