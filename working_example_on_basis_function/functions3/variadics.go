package functions

import "fmt"

func CreateVariadic(users ...string) {

	/*

		Variadic functions can be called with any number of trailing arguments.
		For example, fmt.Println is a common variadic function.

	*/

	created_users := []string{}

	for ind := 0; ind < len(users); ind++ {

		fmt.Printf("User %v created\n", users[ind])
		created_users = append(created_users, users[ind])
	}

	fmt.Println("All created users are ", created_users)

}
