package main

import "fmt"

type Person struct {
	//properties
	Name    string
	Surname string
	Age     int
}

func main() {
	p := Person{"era", "asddada", 34}
	fmt.Println(p)
	p2 := Person{Age: 34, Surname: "123", Name: "2323"}
	fmt.Println(p2)
}
