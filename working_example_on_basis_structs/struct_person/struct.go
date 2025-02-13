package main

import "fmt"

type Person struct {
	Id       string
	Fullname string
	Email    string
}

func (p Person) String() string {
	return fmt.Sprintf("%s : %s : %s", p.Id, p.Fullname, p.Email)
}

func main() {
	type personrequest struct {
		FullName string
		Email    string
	}

	pRequest := personrequest{
		FullName: "Jack London",
		Email:    "chinhtt1@onemount.com",
	}

	person := Person{
		Id:       "ot-68",
		Fullname: pRequest.FullName,
		Email:    pRequest.Email,
	}

	chinhinfo := Person{
		Id:       "xxxxxxx",
		Fullname: "Garri Dain",
		Email:    "tranthechinh@gmail.com",
	}

	fmt.Println(person)
	fmt.Println(chinhinfo)
}


// ot-68 : Jack London : chinhtt1@onemount.com
// xxxxxxx : Garri Dain : tranthechinh@gmail.com
