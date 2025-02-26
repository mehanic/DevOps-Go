package main

import "fmt"

// Assigns an anonymous function to a variable sayHelloWorld
var sayHelloWorld = func() {
	fmt.Println("Hello World !")
}

func main() {
	sayHelloWorld() // Hello World !
}

