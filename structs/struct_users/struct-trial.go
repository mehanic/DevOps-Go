package main

import "fmt"

var userData [2]Users
var user1 Users
var user2 Users

func main() {
	user1 = Users{"1", "tiq", "2412"}
	user2 = Users{"2", "design", "3461"}
	userData = [2]Users{user1, user2}
	
	fmt.Println("------------------------------------")
	printUsers()
}

type Users struct {
	userId   string
	username string
	password string
}

func printUsers() {
	for i := 0; i < len(userData); i++ {
		fmt.Println("name surname:", userData[i].username)
		fmt.Println("numbers:", userData[i].password)
		fmt.Println("------------------------------------")
	}
}


// ------------------------------------
// name surname: tiq
// numbers: 2412
// ------------------------------------
// name surname: design
// numbers: 3461
// ------------------------------------
