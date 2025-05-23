package main

import "fmt"

func showFullName(s Student) {
	fmt.Println(s.FirstName, s.LastName, s.Region)
}

type Student struct {
	FirstName string
	LastName  string
	Region    string
}

func main() {
	s1 := Student{
		FirstName: "Peter",
		LastName:  "Lviv",
		Region:    "Praha",
	}
	s2 := Student{
		"Dmitro",
		"Cherson",
		"Dresden",
	}
	students := []Student{s1, s2}
	for _, v := range students {
		showFullName(v)
	}

}

// Peter Lviv Praha
// Dmitro Cherson Dresden
