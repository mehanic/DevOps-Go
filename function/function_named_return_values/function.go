package main

import (
	"fmt"
)


func main() {
	firstName, middleName, lastName := fullName()
	fmt.Println(firstName, middleName, lastName)
}

func fullName() (firstName, middleName, lastName string) {
	firstName = "first"
	middleName = "middele"
	lastName = "last"
	return
}