package main

import "fmt"

// Hello is an exported function
func Hello() {
	fmt.Println("exported hello")
}

func main() {
	Hello()
}

// exported hello
