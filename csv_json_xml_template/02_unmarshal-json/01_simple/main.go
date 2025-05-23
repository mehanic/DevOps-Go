package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	bs := []byte(`{"Name":"CLE","Age":29}`)
	var p Person
	err := json.Unmarshal(bs, &p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)
}

//{CLE 29}
