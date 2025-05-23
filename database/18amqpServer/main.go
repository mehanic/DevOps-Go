package main

import (
	"github.com/djumanoff/amqp"
	"github.com/kirigaikabuto/Golang1300Lessons/18amqpServer/utils"
)

func main() {
	config := amqp.Config{
		Host: "localhost",
		Port: 5672,
	}
	serverConfig := amqp.ServerConfig{
		ResponseX: "response",
		RequestX:  "request",
	}
	sess := amqp.NewSession(config)
	err := sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	srv, err := sess.Server(serverConfig)
	if err != nil {
		panic(err)
		return
	}
	srv.Endpoint("lesson18.GetNumber", utils.AMQPGetNumber())
	srv.Endpoint("lesson18.GetProducts", utils.AMQPGetProducts())
	if err := srv.Start(); err != nil {
		panic(err)
		return
	}
}
