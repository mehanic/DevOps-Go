package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type student struct {
	person
	id int
}

func (p person) Greeting() {
	fmt.Println("Congratulation person")
}

func (s student) Greeting() {
	fmt.Println("Congratulation student")
}

func main() {

	p := person{firstName: "Clement", lastName: "Tsai", age: 29}
	p.Greeting()

	s := student{
		person: person{"clem", "tsai", 30},
		id:     1,
	}
	s.Greeting()
	s.person.Greeting()
}

// Congratulation person
// Congratulation student
// Congratulation person
