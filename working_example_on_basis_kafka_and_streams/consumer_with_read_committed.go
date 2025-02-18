package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Create a consumer with read_committed isolation
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-consumer-group",
		"isolation.level":   "read_committed",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to the transactional topic
	topic := "transactional-topic"
	if err := consumer.Subscribe(topic, nil); err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}

	// Consume and print messages
	fmt.Println("Waiting for committed messages...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received committed message: %s\n", string(msg.Value))
		} else {
			log.Printf("Consumer error: %v\n", err)
		}
	}
}

