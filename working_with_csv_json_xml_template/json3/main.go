package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Student struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}
	type Class struct {
		Name     string    `json:"class_name"`
		Students []Student `json:"students"`
	}

	//studentOne := Student{Name: "Aaron", Age: 24, Gender: "Male"}
	studentTwo := Student{Name: "Jane", Age: 24, Gender: "Female"}

	students := []Student{{"Charles", 12, "male"}, studentTwo}

	class := Class{"JS1", students}

	classJson, err := json.Marshal(class)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(classJson))

}

//{"class_name":"JS1","students":[{"name":"Charles","age":12,"gender":"male"},{"name":"Jane","age":24,"gender":"Female"}]}
