package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	p := Person{"Bapan", 21}
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error occured:")
		return
	}
	fmt.Println(string(data))
	jsaondata := json.RawMessage(string(data))
	bte, err := jsaondata.MarshalJSON()
	if err != nil {
		fmt.Println(err)
		return
	}
	p1 := Person{}
	err = json.Unmarshal(bte, &p1)
	if err != nil {
		fmt.Println("Error occured")
		return
	}
	fmt.Println(p1)
}
