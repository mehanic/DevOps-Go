package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "transactional-consumer",
		"auto.offset.reset": "earliest",
		"isolation.level":   "read_committed",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to the transactions topic
	if err := consumer.SubscribeTopics([]string{"transactions"}, nil); err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	fmt.Println("Consumer started with read_committed isolation. Waiting for messages...")

	// Consume messages
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v\n", err)
		}
	}
}


//kafka-console-consumer.sh --bootstrap-server localhost:9092 \
//  --topic transactions \
//  --from-beginning \
//  --property print.key=true \
//  --isolation-level read_committed


//If the transaction was committed → Messages appear.
//If aborted → No messages appear due to read_committed.
