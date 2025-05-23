package main

import "fmt"

type Contact struct {
	Phone, Email string
}

func main() {
	myContact := Contact{Phone: "898989", Email: "airondev@gmail.com"}
	var yourContact Contact = Contact{"90909", "airondev@gmail.com"}
	theirContact := &Contact{"989898", "iiuiu"}

	fmt.Println("myContact", myContact)
	fmt.Println("yourContact", yourContact)
	fmt.Println("theirContact", *theirContact)

}

// myContact {898989 airondev@gmail.com}
// yourContact {90909 airondev@gmail.com}
// theirContact {989898 iiuiu}

