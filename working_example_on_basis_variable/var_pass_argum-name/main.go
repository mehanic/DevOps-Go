package main

import "fmt"

func main() {
	var greeting string
	fmt.Print("input your name: ")
	fmt.Scanln(&greeting)
	fmt.Println("welcome to our company, " + greeting + "!")
}
