package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func InputUserData() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser := createNewUser(firstName, lastName, birthDate)

	newUser.outputUserData()
	
}

func (user *User) outputUserData() {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')

	var cleanedInput string

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}

	return cleanedInput
}

type User struct {
	firstName, lastName, birthDate string
	created time.Time
}

func createNewUser(fName, lName, bDate string) *User {
	return &User{
		fName,
		lName,
		bDate,
		time.Now(),
	}
}

func main() {
	// Call the InputUserData function from the structures package
	InputUserData()
}

// Please enter your first name: peter
// Please enter your last name: new
// Please enter your birthdate (MM/DD/YYYY): 01/01/2020
// My name is peter new (born on 01/01/2020)
