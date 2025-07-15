package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	topic := "logs"

	conn, err := kafka.DialContext(ctx, "tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
	})
	if err != nil {
		panic(err)
	}
}