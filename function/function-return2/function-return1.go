package main

import "fmt"

func add(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	fullName := firstName + " " + lastName
	return z, fullName
}

func main() {
	age, name := add(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson
}
