package main

import (
	"fmt"
)

var _ User = &superUser{}

type superUser struct {
	Name      string
	Age       int
	isBlocked bool
}

func (u *superUser) Block() {
	u.isBlocked = true
	fmt.Printf("superUser %s is now blocked.\n", u.Name)
}

var _ User = &user{}

type user struct {
	Adress, Fio, Phone string
	isBlocked          bool
}

func (u *user) Block() {
	u.isBlocked = true
	fmt.Printf("User %s with address %s is now blocked.\n", u.Fio, u.Adress)
}

type User interface {
	Block()
}

func NewUser(fio, adress, phone string) User {
	u := user{
		Fio:    fio,
		Adress: adress,
		Phone:  phone,
	}
	return &u
}

func main() {
	us := NewUser("John Doe", "123 Main St", "555-1234")
	us.Block()
}

//User John Doe with address 123 Main St is now blocked.