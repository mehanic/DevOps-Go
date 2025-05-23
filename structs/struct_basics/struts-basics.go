package main

import "fmt"

func main() {
	/*
		USING VARIABLES
	*/

	var (
		name, lastname string
		age            int

		name2, lastname2 string
		age2             int
	)

	name, lastname, age = "Pablo", "Picasso", 95
	name2, lastname2, age = "Sigmund", "Freud", 83

	fmt.Println("Picasso:", name, lastname, age)
	fmt.Println("Freud  :", name2, lastname2, age2)

	// var picasso struct {
	// 	name, lastname string
	// 	age            int
	// }

	// var freud struct {
	// 	name, lastname string
	// 	age            int
	// }

	// create a new struct type
	type person struct {
		name, lastname string
		age            int
	}

	// picasso := person{name: "Pablo", lastname: "Picasso", age: 91}
	picasso := person{
		name:     "Pablo",
		lastname: "Picasso",
		age:      91,
	}

	var freud person
	freud.name = "Sigmund"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("\n%s's age is %d\n", picasso.lastname, picasso.age)

	fmt.Printf("\nPicasso: %#v\n", picasso)
	fmt.Printf("Freud  : %#v\n", freud)
}

//
//Picasso: Pablo Picasso 83
//Freud  : Sigmund Freud 0
//
//Picasso's age is 91
//
//Picasso: main.person{name:"Pablo", lastname:"Picasso", age:91}
//Freud  : main.person{name:"Sigmund", lastname:"Freud", age:83}
//

//----

type person0 struct {
	name string
	age  int
}

func main1() {

	var tom = person0{name: "Tom", age: 24}
	fmt.Println(tom.name) // Tom
	fmt.Println(tom.age)  // 24

	tom.age = 38                   
	fmt.Println(tom.name, tom.age) // Tom 38
}

//----

// Picasso: Pablo Picasso 83
// Freud  : Sigmund Freud 0

// Picasso's age is 91

// Picasso: main.person{name:"Pablo", lastname:"Picasso", age:91}
// Freud  : main.person{name:"Sigmund", lastname:"Freud", age:83}
