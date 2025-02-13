package main

import (
	"encoding/json"
	"fmt"
)

type mydata struct {
	FirstName  string `json:"firstname"`
	SecondName string `json:"-"`
	Age        int    `json:"myage"`
}

func main() {
	x := mydata{"shine", "sivadasan", 31}
	y, _ := json.Marshal(x)
	fmt.Println(string(y))
}

//The SecondName field will be ignored in json since we put tag to ignored
//The other two field names will be changed to names specified in the json tag

//{"firstname":"shine","myage":31}
