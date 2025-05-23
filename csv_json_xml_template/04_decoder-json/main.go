package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	reader := strings.NewReader(`{"Name":"CLE","Age":29}`)
	p := Person{}
	json.NewDecoder(reader).Decode(&p)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}

//CLE
//29