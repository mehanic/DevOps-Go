package main

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
	"github.com/kirigaikabuto/lesson20acl"
)

func main() {
	amqpConfig := amqp.Config{
		Host: "localhost",
		Port: 5672,
	}
	sess := amqp.NewSession(amqpConfig)
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
	cmd := &lesson20acl.CreateRoleCommand{Name: "123456"}
	jsonCmd, err := json.Marshal(cmd)
	if err != nil {
		panic(err)
		return
	}
	jsonData, err := clt.Call("roles.create", amqp.Message{Body: jsonCmd})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(jsonData.Body))
}
