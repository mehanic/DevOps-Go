package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

func main() {
	var p person = person{
		name: "clement",
		age:  29,
		sex:  "male",
	}
	fmt.Println(p)
}

//{clement 29 male}