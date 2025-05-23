package main

import "fmt"

func main() {
	//Assigns an anonymous function to a variable sayHelloWorld
	sayHelloWorld := func(userName string) {
		fmt.Println("Hello", userName)
	}
	//Invoke  sayHelloWorld
	sayHelloWorld("Sharad") // Print : Hello Sharad
}

