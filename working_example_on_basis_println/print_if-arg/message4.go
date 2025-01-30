package main

import (
	"fmt"
//	"os"
//	"strconv"
//	"strings"
)

func main() {
	// Reading input from standard input
	var userStart int
	var userEnd int
	var branch string

	fmt.Print("Enter start user number: ")
	_, err := fmt.Scanf("%d", &userStart)
	if err != nil {
		fmt.Println("Error reading user start:", err)
		return
	}

	fmt.Print("Enter end user number: ")
	_, err = fmt.Scanf("%d", &userEnd)
	if err != nil {
		fmt.Println("Error reading user end:", err)
		return
	}

	fmt.Print("Enter branch name: ")
	_, err = fmt.Scanf("%s", &branch)
	if err != nil {
		fmt.Println("Error reading branch:", err)
		return
	}

	// Generating and printing user IDs and passwords
	for i := userStart; i < userEnd; i++ {
		user := fmt.Sprintf("%03d", i)
		userID := fmt.Sprintf("user_id is 16/%s/%s", branch, user)
		password := fmt.Sprintf("password is 16/%s/%s", branch, user)

		fmt.Println(userID)
		fmt.Println(password)
	}
}

// Enter start user number: 34
// Enter end user number: 67
// Enter branch name: red
// user_id is 16/red/034
// password is 16/red/034
// user_id is 16/red/035
// password is 16/red/035
// user_id is 16/red/036
