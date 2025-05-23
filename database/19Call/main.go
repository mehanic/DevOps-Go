package main

import (
	"encoding/json"
	"fmt"
	"github.com/djumanoff/amqp"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

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
	//get user by id
	getUserRequest := &User{Id: "0b18e398-7928-4cbd-b669-6ba310e7692a"}
	getUserRequestJson, err := json.Marshal(getUserRequest)
	if err != nil {
		panic(err)
		return
	}
	response, err := clt.Call("users.getById", amqp.Message{Body: getUserRequestJson})
	if err != nil {
		panic(err)
		return
	}
	getUserResponse := &User{}
	err = json.Unmarshal(response.Body, &getUserResponse)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(getUserResponse)
	//create user
	//createUserRequest := &User{Username: "1234", Password: "1234"}
	//createUserRequestJson, err := json.Marshal(createUserRequest)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//response, err := clt.Call("users.create", amqp.Message{Body: createUserRequestJson})
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//createUserResponse := &User{}
	//err = json.Unmarshal(response.Body, &createUserResponse)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Println(createUserResponse)
}
