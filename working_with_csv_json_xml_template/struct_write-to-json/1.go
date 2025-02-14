package main

import (
	"encoding/json"
	"fmt"

	"os"
)

// struct
type Student struct {
	FirstName string
	LastName  string
	Age       int
	FullName  string
	YearBirth int
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

func main() {
	var st1 Student
	st1 = Student{
		FirstName: "Peter",
		LastName:  "Praha",
		Age:       23,
	}
	st2 := Student{
		FirstName: "Nikol",
		LastName:  "Wien",
		Age:       123,
	}
	students := []Student{st1, st2}
	fmt.Println(students)
	//convert to json
	studentsJson, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}
	//write data to file
	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(studentsJson)
	if err != nil {
		panic(err)
	}
}

 
//[{Peter Praha 23  0} {Nikol Wien 123  0}]
