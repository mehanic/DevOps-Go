package main

import "fmt"

type contact struct {
	email string
	phone string
}

type person struct {
	name        string
	age         int
	contactInfo contact
}

func main() {

	var tom = person{
		name: "Tom",
		age:  24,
		contactInfo: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	tom.contactInfo.email = "supertom@gmail.com"

	fmt.Println(tom.contactInfo.email) // supertom@gmail.com
	fmt.Println(tom.contactInfo.phone) // +1234567899
}

