package main

import (
	"fmt"
)

type mydata struct {
	firstname  string
	secondname string
}

func main() {
	p := mydata{"shine", "sivadasan"}
	fmt.Println(p.firstname, p.secondname)

	x := mydata{firstname: "gargi", secondname: "sasidharan"}
	fmt.Printf(x.firstname, "\t", x.secondname)
}

// shine sivadasan
// gargi%!(EXTRA string=	, string=sasidharan)
