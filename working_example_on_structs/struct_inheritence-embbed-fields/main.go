package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person        // embbed type, inheritence
	name   string // override
	id     int
	class  string
}

func main() {

	st1 := student{
		person: person{
			name: "Cle",
			age:  29,
		},
		name:  "Outer",
		id:    1,
		class: "A",
	}

	st2 := student{
		person: person{
			name: "John",
			age:  20,
		},
		id:    2,
		class: "B",
	}

	fmt.Println(st1)
	// fields and method of the inner type is promoted to the outer type
	fmt.Println(st1.name, st1.person.name)
	fmt.Println(st1.age, st1.person.age)

	fmt.Println(st2)
	fmt.Println(st2.name, st2.person.name) // st2.name is empty string
	fmt.Println(st2.age, st2.person.age)
}


// {{Cle 29} Outer 1 A}
// Outer Cle
// 29 29
// {{John 20}  2 B}
//  John
// 20 20
