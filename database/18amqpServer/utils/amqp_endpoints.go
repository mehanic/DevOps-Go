package utils

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

func AMQPGetNumber() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		res := GetNumber()
		response := &amqp.Message{Body: []byte(res)}
		return response
	}
}

func AMQPGetProducts() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		products := GetProducts()
		dataJson, _ := json.Marshal(products)
		return &amqp.Message{Body: dataJson}
	}
}
