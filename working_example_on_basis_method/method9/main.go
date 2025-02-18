package main

import "fmt"

func main() {
	fmt.Println("This is a class on methods in GoLang. We are working with what we used in learning structs")

	mayowa := User{"Mayowa", "mayowa@go.dev", true, 23}
	fmt.Println(mayowa)
	fmt.Printf("Mayowa's details are: %+v\n", mayowa)
	fmt.Printf("Name is %v and email is %v.\n", mayowa.Name, mayowa.Email)

	mayowa.GetStatus()
	mayowa.NewMail()

	fmt.Printf("Name is %v and email is %v.\n", mayowa.Name, mayowa.Email)
}

// Note: No inheritance in golang. Also, no super or parent

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

// This is a class on methods in GoLang. We are working with what we used in learning structs
// {Mayowa mayowa@go.dev true 23}
// Mayowa's details are: {Name:Mayowa Email:mayowa@go.dev Status:true Age:23}
// Name is Mayowa and email is mayowa@go.dev.
// Is user active:  true
// Email of this user is:  test@go.dev
// Name is Mayowa and email is mayowa@go.dev.

