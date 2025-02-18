package structs

import "fmt"

func CreateStruct() {

	/*

		A struct (short for structure) is used to create a collection of members of different data types,
		into a single variable.

		While arrays are used to store multiple values of the same data type into a single variable,
		structs are used to store multiple values of different data types into a single variable.

		A struct can be useful for grouping data together to create records.

	*/

	// To declare a structure in Go, use the type and struct keywords.

	type User struct {
		username    string
		mail        string
		birthyear   int
		job         string
		phonenumber string
	}

	var user1 User // This is how we create a variable

	user1.username = "berkayalan"
	user1.mail = "berkayalan.mail@gmail.com"
	user1.birthyear = 1996
	user1.job = "Cloud Engineer"
	user1.phonenumber = "+44 123456789"

	// or we can create like user1 := User{"berkayalan","berkayalan.mail@gmail.com",1996,"Cloud Engineer","+44 123456789"}

	fmt.Println(user1)

}
