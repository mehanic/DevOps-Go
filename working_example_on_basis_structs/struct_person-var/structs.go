package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	var acc Account = Account{
		Id:   1,
		Name: "Hoasker",
		Person: Person{
			Name:    "RealZealot",
			Address: "Qazaqstan",
		},
	}
	fmt.Println("%#v\n", acc)

	acc.Owner = Person{2, "Hoasker", "KZ"}

	fmt.Printf("%#v\n", acc)

	fmt.Println(acc.Address)
	fmt.Println(acc.Person.Name)
}

// %#v
//  {1 Hoasker <nil> {0  } {0 RealZealot Qazaqstan}}
// main.Account{Id:1, Name:"Hoasker", Cleaner:(func(string) string)(nil), Owner:main.Person{Id:2, Name:"Hoasker", Address:"KZ"}, Person:main.Person{Id:0, Name:"RealZealot", Address:"Qazaqstan"}}
// Qazaqstan
// RealZealot
