package main

import (
	"fmt"
)


type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	
	var p Person
	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)
	fmt.Println(p.Age)
	p.FirstName = "asddassa"
	p.LastName = "12321313123"
	p.Age = 30
	fmt.Println(p.FirstName)
	fmt.Println(p.LastName)
	fmt.Println(p.Age)
}

// 0
// asddassa
// 12321313123
// 30
