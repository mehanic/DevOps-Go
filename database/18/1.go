package main

import (
	"encoding/json"
	"github.com/kirigaikabuto/Golang1300Lessons/18/users"
	"github.com/streadway/amqp"
)

func main() {
	//producer -> publisher
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
		return
	}
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
		return
	}
	queue, err := channel.QueueDeclare(
		"users18",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
		return
	}
	user1 := &users.User{Id: "yerassyl", Name: "tleugazy"}
	jsonData, err := json.Marshal(user1)
	if err != nil {
		panic(err)
		return
	}
	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonData,
		},
	)
	if err != nil {
		panic(err)
		return
	}

}
