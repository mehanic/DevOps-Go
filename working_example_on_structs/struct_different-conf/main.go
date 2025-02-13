package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}


func main() {

	e1 := employee{
		name:      "Gurcan",
		age:       40,
		isMarried: true,
	}

	fmt.Println(e1)


	m1 := manager{}
	m1.name = "Ayşe"
	m1.age = 28
	m1.isMarried = false
	m1.hasDegree = true

	fmt.Println(m1.employee)

	// Anonim Struct

	theBoss := struct {
		name  string
		money bool
	}{name: "THE BOSS", money: true}

	fmt.Println(theBoss)
}

// {Gurcan 40 true}
// {Ayşe 28 false}
// {THE BOSS true}
