package main

import (
	"fmt"
)

type mydata struct {
	firstname  string
	secondname string
}

func (data mydata) concatenate() string {
	return data.firstname + " " + data.secondname
}


func main() {
	x := mydata{"shine", "sivadasan"}
	fmt.Println(x.concatenate())
}

//shine sivadasan