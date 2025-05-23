package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Create a transactional producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    "localhost:9092",
		"acks":                 "all",
		"enable.idempotence":   true,
		"transactional.id":     "my-transactional-producer",
		"retries":              5,
		"max.in.flight.requests.per.connection": 5,
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	// Initialize transactions
	if err := producer.InitTransactions(nil); err != nil {
		log.Fatalf("Failed to initialize transactions: %v", err)
	}

	// Begin the transaction
	if err := producer.BeginTransaction(); err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}

	// Produce messages within the transaction
	topic := "transactional-topic"
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Transactional Message %d", i)
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
			producer.AbortTransaction(nil)
			return
		}
	}

	// Commit the transaction
	if err := producer.CommitTransaction(nil); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Println("Transaction committed successfully.")
}

