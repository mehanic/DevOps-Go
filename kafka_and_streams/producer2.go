package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// Create producer with transactional.id
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"transactional.id":  "my-transactional-producer",
		"enable.idempotence": true,
		"acks":               "all",
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	// Initialize transactions
	if err := producer.InitTransactions(nil); err != nil {
		log.Fatalf("Failed to init transactions: %v", err)
	}

	// Start transaction
	if err := producer.BeginTransaction(); err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}

	// Produce messages
	topic := "transactions"
	for i := 1; i <= 2; i++ {
		msg := fmt.Sprintf("Order #%d", i)
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, nil)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		} else {
			fmt.Printf("Produced: %s\n", msg)
		}
	}

	// Decide whether to commit or abort based on argument
	if len(os.Args) > 1 && os.Args[1] == "abort" {
		// Abort transaction
		if err := producer.AbortTransaction(nil); err != nil {
			log.Fatalf("Failed to abort transaction: %v", err)
		}
		fmt.Println("Transaction aborted.")
	} else {
		// Commit transaction
		if err := producer.CommitTransaction(nil); err != nil {
			log.Fatalf("Failed to commit transaction: %v", err)
		}
		fmt.Println("Transaction committed.")
	}

	// Handle shutdown gracefully
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	<-sigchan
}

