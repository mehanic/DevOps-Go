package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/Golang1300Lessons/18/users"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//consumer
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
	messages, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)

	go func() {
		for d := range messages {
			u := &users.User{}
			err = json.Unmarshal(d.Body, &u)
			if err != nil {
				panic(err)
				return
			}
			fmt.Println(u.Id, u.Name)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
