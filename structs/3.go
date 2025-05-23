package main

import "fmt"

//class
type Person struct {
	//properties
	Name    string
	Surname string
	Age     int
}

func main() {
	var p1 Person
	p1.Name = "Yerassyl"
	p1.Age = 23
	p1.Surname = "Tleugazy"
	fmt.Println(p1.Name)
	fmt.Println(p1.Surname)
	fmt.Println(p1.Age)
}
