package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a new struct for the top-level object
type Students struct {
	Students []Student `json:"students"`
}

type Student struct {
	FirstName string
	LastName  string
	Age       int
	FullName  string
	YearBirth int
}

func main() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	// Declare the top-level struct, not just []Student
	var studentsList Students
	err = json.Unmarshal(data, &studentsList)
	if err != nil {
		panic(err)
	}

	// Loop through and print each student's first name
	for _, v := range studentsList.Students {
		fmt.Println(v.FirstName)
	}
}

// methods
func (st *Student) PrintAll() {
	st.CreateFullName()
	st.SetYear()
	fmt.Println(st.FirstName, st.LastName, st.Age, st.FullName, st.YearBirth)
}

func (st *Student) CreateFullName() {
	st.FullName = st.FirstName + " " + st.LastName
}

func (st *Student) SetYear() {
	st.YearBirth = 2021 - st.Age
}
