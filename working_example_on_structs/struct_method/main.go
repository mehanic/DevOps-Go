package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) fullName() string {
	return p.firstName + p.lastName
}

func main() {

	p := person{firstName: "Clement", lastName: "Tsai", age: 29}

	fmt.Println(p.fullName())
}

//ClementTsai