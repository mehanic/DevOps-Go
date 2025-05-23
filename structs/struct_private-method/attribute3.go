package main

import (
	"fmt"
)

// Define the Person struct
type Person struct {
	name string
	age  int
}

// Method to assign name and age and display it
func (p *Person) AssignNameAndAge(name string, age int) {
	p.name = name
	p.age = age
	p.display()
}

// Private method to display name and age
func (p *Person) display() {
	fmt.Println(p.name, p.age)
}

func main() {
	per1 := Person{}
	per1.AssignNameAndAge("John", 32)
}

//John 32
