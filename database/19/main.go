package main

import (
	"github.com/djumanoff/amqp"
	"github.com/kirigaikabuto/users19"
)

func main() {
	mongoUsersStore, err := users.NewUsersStore(users.MongoConfig{
		Host:           "localhost",
		Port:           "27017",
		Database:       "lesson19",
		CollectionName: "users",
	})
	if err != nil {
		panic(err)
		return
	}
	amqpConfig := amqp.Config{
		Host: "localhost",
		Port: 5672,
	}
	serverConfig := amqp.ServerConfig{
		ResponseX: "response",
		RequestX:  "request",
	}
	sess := amqp.NewSession(amqpConfig)
	err = sess.Connect()
	if err != nil {
		panic(err)
		return
	}
	srv, err := sess.Server(serverConfig)
	if err != nil {
		panic(err)
		return
	}
	amqpEndpoints := users.NewUsersAmqpEndpoints(mongoUsersStore)
	srv.Endpoint("users.create", amqpEndpoints.CreateUserAmqpEndpoint())
	srv.Endpoint("users.getById", amqpEndpoints.GetUserAmqpEndpoint())
	if err := srv.Start(); err != nil {
		panic(err)
		return
	}

}
