package interface_module

import (
	"fmt"
)

/*

	An interface type is defined as a set of method signatures.

	A value of interface type can hold any value that implements those methods.

	An interface defines a behavior of a type.

*/

type User interface {
	CreateUser() string
	DeleteUser() string
}

type userinfo struct {
	username  string
	email     string
	birthyear int
}

func (u userinfo) CreateUser() string {
	return fmt.Sprintf("User Created : %v", u.username)
}

func (u userinfo) DeleteUser() string {
	return fmt.Sprintf("User Deleted : %v", u.username)
}

func CreateInterface(u User) {
	fmt.Println(u.CreateUser())

	fmt.Println(u.DeleteUser())
}

func UserInterface() {

	newuser := userinfo{"berkayalan", "berkayalan.mail@gmail.com", 1996}

	CreateInterface(newuser)

}
