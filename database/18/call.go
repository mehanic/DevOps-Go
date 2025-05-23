package main

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
)

type Product struct {
	Id    int
	Name  string
	price float64
}

func main() {
	config := amqp.Config{
		Host: "localhost",
		Port: 5672,
	}
	sess := amqp.NewSession(config)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	clt, err := sess.Client(amqp.ClientConfig{})
	if err != nil {
		panic(err)
		return
	}
	res, err := clt.Call("lesson18.GetNumber", amqp.Message{})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(res.Body))
	res, err = clt.Call("lesson18.GetProducts", amqp.Message{})
	if err != nil {
		panic(err)
		return
	}
	prs := []Product{}
	err = json.Unmarshal(res.Body, &prs)
	if err != nil {
		panic(err)
		return
	}
	for i, v := range prs {
		fmt.Println(i, v.Id, v.Name, v.price)
	}
}
