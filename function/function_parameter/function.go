package main

import "fmt"

func main() {
	fullName("Ivan", "Mazepa")
}

func fullName(firstName string, lastName string) {
	fmt.Println(firstName + " " + lastName)
}
