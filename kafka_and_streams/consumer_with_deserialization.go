package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Order struct matches the producer struct
type Order struct {
	OrderID string `json:"order_id"`
	Amount  int    `json:"amount"`
}

func main() {
	// Create Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "serialization-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	topic := "serialization-demo"
	consumer.Subscribe(topic, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v", err)
			continue
		}

		// Deserialize as string
		message := string(msg.Value)
		fmt.Printf("Raw message: %s\n", message)

		// Try to deserialize as JSON
		var order Order
		if err := json.Unmarshal(msg.Value, &order); err == nil {
			fmt.Printf("Deserialized JSON Order: ID=%s, Amount=%d\n", order.OrderID, order.Amount)
		}
	}
}

