package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Order struct for JSON serialization
type Order struct {
	OrderID string `json:"order_id"`
	Amount  int    `json:"amount"`
}

func main() {
	// Create Kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	topic := "serialization-demo"

	// Example 1: String serialization
	stringMessage := "Hello, Kafka!"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(stringMessage),
	}, nil)

	fmt.Println("Produced string message:", stringMessage)

	// Example 2: JSON serialization
	order := Order{OrderID: "1234", Amount: 500}
	jsonData, _ := json.Marshal(order)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          jsonData,
	}, nil)

	fmt.Println("Produced JSON message:", string(orderJSON))

	// Wait for message delivery
	producer.Flush(1000)
}

